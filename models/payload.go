package models

type FactorialRequest struct {
	AFactorial *int `json:"a"`
	BFactorial *int `json:"b"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
