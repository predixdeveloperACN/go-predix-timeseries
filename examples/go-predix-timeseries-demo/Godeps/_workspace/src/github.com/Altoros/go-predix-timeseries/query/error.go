package query

import (
	"encoding/json"
	"fmt"
)

type QueryError struct {
	Errors        []string `json:"errors"`
	CorrelationID string   `json:"correlationID"`
}

func Error(data []byte) *QueryError {
	var res QueryError
	err := json.Unmarshal(data, &res)
	if err != nil {
		res.Errors = append(res.Errors, fmt.Sprintf("Query error unmarshaling failed: %s", err))
		res.CorrelationID = "Not Applicable"
	}
	return &res
}

func (qe *QueryError) Error() string {
	return fmt.Sprintf("Query errors: %q, Correlation ID: %s", qe.Errors, qe.CorrelationID)
}
