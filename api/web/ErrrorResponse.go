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

func NewErrorsReponse(errs []error) ErrorResponse {

	errors := make([]string, len(errs))
	for index, err := range errs {
		errors[index] = err.Error()
	}

	return ErrorResponse{
		Errors: errors,
	}
}

func NewInternalError() ErrorResponse {
	errors := make([]string, 1)
	errors[0] = "Internal server error"
	return ErrorResponse{
		Errors: errors,
	}
}
