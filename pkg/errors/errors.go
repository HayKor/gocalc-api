package errors

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func New(statusCode int, message string) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e APIError) Error() string {
	return e.Message
}
