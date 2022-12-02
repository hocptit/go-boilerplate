package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	//todo: hidden when prod
	Errors     any    `json:"errors"`
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Success    bool   `json:"success"`
	MetaData   any    `json:"metaData"`
}

func ReturnData(c *gin.Context, code int, data any) {
	response := Response{
		Data:       data,
		Message:    "",
		Errors:     "",
		ErrorCode:  "",
		StatusCode: code,
		Success:    true,
	}
	c.JSON(http.StatusOK, response)
	return
}
