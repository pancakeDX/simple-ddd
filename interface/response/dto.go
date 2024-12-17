package response

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"error"`
}

func NewErrorResponse(errorMessge string) *ErrorResponse {
	return &ErrorResponse{
		ErrorMessage: errorMessge,
	}
}
