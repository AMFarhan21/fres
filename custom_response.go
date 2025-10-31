package fres

type (
	custom struct{}

	SuccessResponse struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}

	ErrorResponse struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Error   interface{} `json:"error,omitempty"`
	}
)

var Custom = custom{}

func (c custom) SuccessResponse() SuccessResponse {
	return SuccessResponse{}
}

func (c custom) ErrorResponse() ErrorResponse {
	return ErrorResponse{}
}
