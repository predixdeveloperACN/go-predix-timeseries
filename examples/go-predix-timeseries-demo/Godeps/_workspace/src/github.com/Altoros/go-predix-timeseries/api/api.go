package api

import (
	"fmt"
	"net/http"

	"github.com/Altoros/go-predix-timeseries/ingest"
	"github.com/Altoros/go-predix-timeseries/query"
	"github.com/gorilla/websocket"
	"gopkg.in/h2non/gentleman.v0"
)

type Api struct {
	ingestUrl    string
	queryUrl     string
	authToken    string
	predixZoneId string
	conn         *websocket.Conn
	client       *gentleman.Client
}

func New(ingestUrl, queryUrl, authToken, predixZoneId string) *Api {
	return &Api{
		ingestUrl:    ingestUrl,
		queryUrl:     queryUrl,
		authToken:    authToken,
		predixZoneId: predixZoneId,
	}
}

func Ingest(ingestUrl, authToken, predixZoneId string) *Api {
	return New(ingestUrl, "", authToken, predixZoneId)
}

func Query(queryUrl, authToken, predixZoneId string) *Api {
	return New("", queryUrl, authToken, predixZoneId)
}

func (a *Api) IngestMessage() *ingest.Message {
	if a.conn == nil {
		if a.ingestUrl == "" {
			return nil
		}
		h := http.Header{}
		h.Add("Authorization", fmt.Sprintf("Bearer %s", a.authToken))
		h.Add("Predix-Zone-Id", a.predixZoneId)
		h.Add("Origin", "http://localhost")
		conn, _, err := websocket.DefaultDialer.Dial(a.ingestUrl, h)
		if err != nil {
			return nil
		}
		a.conn = conn
	}
	return ingest.NewMessage(a.conn)
}

func (a *Api) Query() query.Query {
	if a.client == nil {
		if a.queryUrl == "" {
			return nil
		}
		a.client = gentleman.New()
		a.client.BaseURL(a.queryUrl)
		a.client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", a.authToken))
		a.client.SetHeader("Predix-Zone-Id", a.predixZoneId)
	}
	return query.New(a.client)
}
