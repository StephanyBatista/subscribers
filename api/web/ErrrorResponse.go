package web

type ErrorResponse struct {
	Errors []string
}

func NewErrorReponse(message string) ErrorResponse {
	errors := make([]string, 1)
	errors[0] = message

	return ErrorResponse{
		Errors: errors,
	}
}

func NewErrorsReponse(messages []string) ErrorResponse {
	return ErrorResponse{
		Errors: messages,
	}
}

func NewInternalError() ErrorResponse {
	errors := make([]string, 1)
	errors[0] = "Internal server error"
	return ErrorResponse{
		Errors: errors,
	}
}
