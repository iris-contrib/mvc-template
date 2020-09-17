package models

// CounterResponse is the response for the counter API.
type CounterResponse struct {
	Value uint64 `json:"value"`
}

// ErrorResponse is the http error response.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
