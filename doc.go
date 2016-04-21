/*
A Go wrapper around the Predix Time Series API. See service(https://www.predix.io/docs/#mnlfuvZz) and API(https://www.predix.io/api#!/Time_Series) documentation for more details.

Data Ingestion

A time series uses tags, which are often represent sensors (for example, a temperature sensor). A tag consists of one Tag Name (Sensor) and a set of data points. A data point consists of a timestamp (set automatically), measurement (value), and quality. The time series tag can also include one or more attributes (key/value pairs) as additional options.

Pushing Time Series Data

The following example shows a data ingestion request:

api := api.Ingest("wss://ingest_url", accessToken, predixZoneId)
m := api.IngestMessage()
t, _ := m.AddTag("test_tag")
t.AddDatapoint(measurement.Int(123), dataquality.Good)
t.SetAttribute("key", "value")
m.Send()

Querying Time Series Data

You can use the time series API to list tags and attributes as well as to query data points. You can query data points using a start time, end time, tag names, time range, measurement, and attributes. The following example shows a data query with a limit of 1000 set for the number of data points to return in the query result:

api := api.Query("https://query_url", accessToken, predixZoneId)
q := api.Query().Interval(query.R(5).MinutesAgo(), query.AbsTime(time.Now()))
q.Tags(query.Tag("ALT_SENSOR", "TEMP_SENSOR").
  Limit(1000).
  Aggregation("avg", aggregation.Interval(1).Hour()).
  Filter(tag.Attributes(tag.Attr("host","<host>"), tag.Attr("type", "<type>"))).
  Filter(measurement  .Ge(measurement.Float(23.1))).
  Filter(dataquality.Qualities(dataquality.Bad, dataquality.Good)).
  GroupByAttributes("attributename1", "attributename2"))

Query Properties

Your query must specify a start time that can be either absolute or relative. An end time is optional and can be either absolute or relative. If you do not specify the end time, the query uses the current date and time.

Using Relative Start and End Times in a Query

This example shows a query using a relative start time and an absolute end time.

q := api.Query().Interval(query.R(1).YearsAgo(), query.AbsTime(time.Now()))

Limiting the Data Points Returned by a Query

q := api.Query().Start(query.R(15).MinutesAgo())
q.Tags(q.Tag("ALT_SENSOR").Limit(1000))

Data points in the query result cannot exceed the maximum limit of 200,000. Narrow your query criteria (for example, time window) to return a fewer number of data points.

Specifying the Order of Data Points Returned by a Query

The default order for query results is ascending, and it is ordered by timestamp. You can use the `InvertOrder` method to invert the order in which the data points returned.

q := api.Query().Start(query.R(15).MinutesAgo())
q.Tags(q.Tag("ALT_SENSOR").InvertOrder())

Querying for the Latest Data Point

* Filters are the only supported operation, and they are optional.

* Specifying a start time is optional. However, if you do specify the start time, you must also specify an end time.

* If you do not define a time window, the query retrieves the latest data points from the current time (now). The following shows an example of a query for the latest data point with no time window defined:

q := api.Query()
q.LatestDatapoints(query.Tag("ALT_SENSOR").Filter(measurement.Le(10)))
*/
package go_predix_timeseries
