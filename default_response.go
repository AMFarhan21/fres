package fres

type response struct{}

var Response = response{}

func (r response) StatusOK(data interface{}) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Message: "Request successful",
		Data:    data,
	}
}

func (r response) StatusCreated(data interface{}) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Message: "Resource created successfully",
		Data:    data,
	}
}

func (r response) StatusBadRequest(error interface{}) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: "Invalid request body",
		Error:   error,
	}
}

func (r response) StatusUnauthorized(error interface{}) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: "Unauthorized",
		Error:   error,
	}
}

func (r response) StatusForbidden(error interface{}) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: "You do not have permission to access this resource",
		Error:   error,
	}
}

func (r response) StatusNotFound(error interface{}) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: "Resource not found",
		Error:   error,
	}
}

func (r response) StatusConflict(error interface{}) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: "Conflict",
		Error:   error,
	}
}

func (r response) StatusInternalServerError(error interface{}) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: "Internal server error",
		Error:   error,
	}
}
