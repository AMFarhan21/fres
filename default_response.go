package fres

type response struct{}

var Response = response{}

func (r response) StatusOK(data interface{}) DefaultSuccessResponse {
	return DefaultSuccessResponse{
		Success: true,
		Message: "Request successful",
		Data:    data,
	}
}

func (r response) StatusCreated(data interface{}) DefaultSuccessResponse {
	return DefaultSuccessResponse{
		Success: true,
		Message: "Resource created successfully",
		Data:    data,
	}
}

func (r response) StatusBadRequest(error interface{}) DefaultErrorResponse {
	return DefaultErrorResponse{
		Success: false,
		Message: "Invalid request body",
		Error:   error,
	}
}

func (r response) StatusUnauthorized(error interface{}) DefaultErrorResponse {
	return DefaultErrorResponse{
		Success: false,
		Message: "Unauthorized",
		Error:   error,
	}
}

func (r response) StatusForbidden(error interface{}) DefaultErrorResponse {
	return DefaultErrorResponse{
		Success: false,
		Message: "You do not have permission to access this resource",
		Error:   error,
	}
}

func (r response) StatusNotFound(error interface{}) DefaultErrorResponse {
	return DefaultErrorResponse{
		Success: false,
		Message: "Resource not found",
		Error:   error,
	}
}

func (r response) StatusConflict(error interface{}) DefaultErrorResponse {
	return DefaultErrorResponse{
		Success: false,
		Message: "Conflict",
		Error:   error,
	}
}

func (r response) StatusInternalServerError(error interface{}) DefaultErrorResponse {
	return DefaultErrorResponse{
		Success: false,
		Message: "Internal server error",
		Error:   error,
	}
}
