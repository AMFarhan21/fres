package response

import (
	"github.com/AMFarhan21/fres/custom"
)

func StatusOK(data interface{}) custom.SuccessResponse {
	return custom.SuccessResponse{
		Success: true,
		Message: "Request successful",
		Data:    data,
	}
}

func StatusCreated(data interface{}) custom.SuccessResponse {
	return custom.SuccessResponse{
		Success: true,
		Message: "Resource created successfully",
		Data:    data,
	}
}

func StatusBadRequest(error interface{}) custom.ErrorResponse {
	return custom.ErrorResponse{
		Success: false,
		Message: "Invalid request body",
		Error:   error,
	}
}

func StatusUnauthorized(error interface{}) custom.ErrorResponse {
	return custom.ErrorResponse{
		Success: false,
		Message: "Unauthorized",
		Error:   error,
	}
}

func StatusForbidden(error interface{}) custom.ErrorResponse {
	return custom.ErrorResponse{
		Success: false,
		Message: "You do not have permission to access this resource",
		Error:   error,
	}
}

func StatusNotFound(error interface{}) custom.ErrorResponse {
	return custom.ErrorResponse{
		Success: false,
		Message: "Resource not found",
		Error:   error,
	}
}

func StatusConflict(error interface{}) custom.ErrorResponse {
	return custom.ErrorResponse{
		Success: false,
		Message: "Conflict",
		Error:   error,
	}
}

func StatusInternalServerError(error interface{}) custom.ErrorResponse {
	return custom.ErrorResponse{
		Success: false,
		Message: "Internal server error",
		Error:   error,
	}
}
