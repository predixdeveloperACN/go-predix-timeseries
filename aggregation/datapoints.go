package aggregation

// The Datapoints function enables to specify the number of data points to return.
func Datapoints(datapoints int) Parameter {
	return Param("sampling", map[string]interface{}{"datapoints": datapoints})
}
