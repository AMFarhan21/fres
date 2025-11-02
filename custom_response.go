package fres

type (
	SuccessResponse struct {
		Success bool        `json:"success"`
		Status  string      `json:"status,omitempty"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}

	ErrorResponse struct {
		Success bool        `json:"success"`
		Status  string      `json:"status,omitempty"`
		Message string      `json:"message,omitempty"`
		Error   interface{} `json:"error,omitempty"`
	}
)
