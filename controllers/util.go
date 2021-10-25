package controllers

type ErrorResponse struct {
	Error string
}

func Error(err string) ErrorResponse {
	return ErrorResponse{
		Error: err,
	}
}
