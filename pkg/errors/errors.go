package errors

type APIError struct {
	StatusCode int
	message    string
}

func New(statusCode int, message string) APIError {
	return APIError{
		StatusCode: statusCode,
		message:    message,
	}
}

func (e APIError) Error() string {
	return e.message
}
