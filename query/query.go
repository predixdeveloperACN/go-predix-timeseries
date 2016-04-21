// Package query provides primitives to specify query parameters and do read requests to Time Series.
package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"gopkg.in/h2non/gentleman.v0"
	"gopkg.in/h2non/gentleman.v0/plugins/body"
)

const (
	datapointsURI       = "/v1/datapoints"
	latestdatapointsURI = "/v1/datapoints/latest"
	tagsURI             = "/v1/tags"
	aggregationsURI     = "/v1/aggregations"
)

type Query interface {
	// Sends a request to Time Series
	Do() (*Result, error)
	// Specifies the start time
	Start(t Time) Query
	// Specifies the end time
	End(t Time) Query
	// Specifies the start and end times
	Interval(start, end Time) Query
	// Adds a tag
	AddTag(t QueryTag)
	// Adds tags
	Tags(t ...QueryTag)
	// Returns the list of all ingested tags
	IngestedTags() ([]string, error)
	// Returns the list of supported aggregations
	SupportedAggregations() ([]string, error)
	// Retrieves the most recent data point
	LatestDatapoints(tag ...QueryTag) (*Result, error)
}

type query struct {
	start, end Time
	tags       []QueryTag
	client     *gentleman.Client
}

func (q *query) MarshalJSON() ([]byte, error) {
	j := make(map[string]interface{})
	j["start"] = q.start.Value()
	if q.end != nil {
		j["end"] = q.end.Value()
	}
	j["tags"] = q.tags

	return json.Marshal(j)
}

func (q *query) Start(t Time) Query {
	q.start = t
	return q
}

func (q *query) End(t Time) Query {
	q.end = t
	return q
}

func (q *query) Interval(start, end Time) Query {
	q.start = start
	q.end = end
	return q
}

func (q *query) Do() (*Result, error) {
	if q.client != nil {
		req := q.client.Post()
		req.Path(datapointsURI)
		req.Use(body.JSON(q))
		res, err := req.Send()
		if err != nil {
			fmt.Printf("Request error: %s\n", err)
			return nil, err
		}
		if res.StatusCode != 200 {
			return nil, Error(res.Bytes())
		}
		var result Result
		if err = json.Unmarshal(res.Bytes(), &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New("No client!")
}

func (q *query) AddTag(t QueryTag) {
	q.tags = append(q.tags, t)
}

func (q *query) Tags(t ...QueryTag) {
	q.tags = append(q.tags, t...)
}

func (q *query) IngestedTags() ([]string, error) {
	//query to GET /v1/tags
	var tags struct {
		Results []string
	}
	res, err := q.client.Get().Path(tagsURI).Do()
	if err == nil {
		if err = json.Unmarshal(res.Bytes(), &tags); err == nil {
			return tags.Results, nil
		}
		log.Printf("%s", res.String())
	}
	return nil, err
}

func (q *query) SupportedAggregations() ([]string, error) {
	var res struct {
		Results []struct {
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"results"`
	}

	resp, err := q.client.Get().Path(aggregationsURI).Do()
	if err == nil {
		if err = json.Unmarshal(resp.Bytes(), &res); err == nil {
			supportedAggregations := make([]string, 0)
			for _, aggregation := range res.Results {
				supportedAggregations = append(supportedAggregations, aggregation.Name)
			}
			return supportedAggregations, nil
		}
	}
	return nil, err
}

func (q *query) LatestDatapoints(tags ...QueryTag) (*Result, error) {
	j := make(map[string]interface{})
	if q.start != nil {
		if q.end == nil {
			return nil, errors.New("If you do specify the start time, you must also specify an end time.")
		}
		j["start"] = q.start
		j["end"] = q.end
	}
	j["tags"] = tags
	if q.client != nil {
		req := q.client.Post()
		req.Path(latestdatapointsURI)
		req.Use(body.JSON(j))
		res, err := req.Send()
		if err != nil {
			fmt.Printf("Request error: %s\n", err)
			return nil, err
		}
		if res.StatusCode != 200 {
			return nil, Error(res.Bytes())
		}
		var result Result
		if err = json.Unmarshal(res.Bytes(), &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New("No client!")
}

func New(c *gentleman.Client) Query {
	return &query{tags: make([]QueryTag, 0), client: c}
}
