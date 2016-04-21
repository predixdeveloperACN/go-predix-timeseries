package aggregation

func Datapoints(datapoints int) Parameter {
	return Param("sampling", map[string]interface{}{"datapoints": datapoints})
}
