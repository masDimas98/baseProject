package utils

import "goTest/internal/model/baseModel"

func SuccessMapper(data any) *baseModel.SuccessResponse {
	return &baseModel.SuccessResponse{
		Status:      "Success",
		SuccessCode: 200,
		Data:        data,
	}
}

func CreatedMapper(data any) *baseModel.SuccessResponse {
	return &baseModel.SuccessResponse{
		Status:      "Success",
		SuccessCode: 201,
		Data:        data,
	}
}

func BadRequestMapper(errorMessage string) *baseModel.ErrorResponse {
	return &baseModel.ErrorResponse{
		Status:       "Bad Request",
		ErrorCode:    400,
		ErrorMessage: errorMessage}
}

func NotFoundMapper(errorMessage string) *baseModel.ErrorResponse {
	return &baseModel.ErrorResponse{
		Status:       "Not Found",
		ErrorCode:    200,
		ErrorMessage: errorMessage}
}

func UserNotFoundMapper(errorMessage string) *baseModel.ErrorResponse {
	return &baseModel.ErrorResponse{
		Status:       "User Not Found",
		ErrorCode:    404,
		ErrorMessage: errorMessage}
}

func UnauthorizedMapper(errorMessage string) *baseModel.ErrorResponse {
	return &baseModel.ErrorResponse{
		Status:       "Unauthorized",
		ErrorCode:    401,
		ErrorMessage: errorMessage}
}

func InternalServerErrorMapper(errorMessage string) *baseModel.ErrorResponse {
	return &baseModel.ErrorResponse{
		Status:       "Internal Server Error",
		ErrorCode:    500,
		ErrorMessage: errorMessage}
}
