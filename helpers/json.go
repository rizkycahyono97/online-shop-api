package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/online-shop-api/model/web"
	"net/http"
)

func JSONInternalErrorResponse(c *gin.Context, message string, err error) {
	if message == "" && err != nil {
		message = err.Error()
	} else if message == "" && err == nil {
		message = "Internal Server Error"
	}

	c.JSON(http.StatusInternalServerError, web.ApiResponse{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: message,
		Data:    nil,
	})
}

func JSONBadRequestResponse(c *gin.Context, message string, err error) {
	if message == "" && err != nil {
		message = err.Error()
	} else if message == "" && err == nil {
		message = "Bad Request"
	}

	c.JSON(http.StatusBadRequest, web.ApiResponse{
		Code:    "BAD_REQUEST",
		Message: message,
		Data:    nil,
	})
}

func JSONForbiddenResponse(c *gin.Context, message string, err error) {
	if message == "" && err != nil {
		message = err.Error()
	} else if message == "" && err == nil {
		message = "Forbidden"
	}

	c.JSON(http.StatusForbidden, web.ApiResponse{
		Code:    "FORBIDDEN",
		Message: message,
		Data:    nil,
	})
}

func JSONNotFoundResponse(c *gin.Context, message string, err error) {
	if message == "" && err != nil {
		message = err.Error()
	} else if message == "" && err == nil {
		message = "Not Found"
	}

	c.JSON(http.StatusNotFound, web.ApiResponse{
		Code:    "NOT_FOUND",
		Message: message,
		Data:    nil,
	})
}

func JSONSuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: message,
		Data:    data,
	})
}
