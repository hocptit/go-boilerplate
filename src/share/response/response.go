package response

import (
	"net/http"
	baseDto "server-go/src/share/base/base_dto"
	"server-go/src/share/constant"

	"github.com/gin-gonic/gin"
)

type PaginationMetaData struct {
	// PageSize     int `json:"pageSize"`
	// Page      int `json:"page"`
	baseDto.PaginationDto
	TotalPage int64 `json:"totalPage"`
}

type WithPagination struct {
	// nolint
	data any
	// nolint
	metaData PaginationMetaData
}

type Response struct {
	// nolint
	Errors any `json:"errors"`
	// nolint
	Data       any    `json:"data"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Success    bool   `json:"success"`
	TraceID    string `json:"traceID"`
	// nolint
	MetaData any `json:"metaData"`
}

// ReturnData
// nolint
func ReturnData(c *gin.Context, code int, data any) {
	var response Response
	response = Response{
		Data:       data,
		Message:    "",
		Errors:     "",
		ErrorCode:  "",
		StatusCode: code,
		Success:    true,
		TraceID:    c.Keys[constant.TraceID].(string),
	}

	c.JSON(http.StatusOK, response)
}

func CalculationTotalPage(totalRecord int64, pageSize int) int64 {
	var add int64
	if (totalRecord % int64(pageSize)) > 0 {
		add = 1
	}
	return (totalRecord / int64(pageSize)) + add
}

// ReturnDataWithPagination
// nolint
func ReturnDataWithPagination(c *gin.Context, data any, totalRecord int64) {
	paginationDto := c.Keys[constant.Pagination].(baseDto.PaginationDto)
	var paginationMetaData = PaginationMetaData{
		TotalPage:     CalculationTotalPage(totalRecord, paginationDto.Limit),
		PaginationDto: paginationDto,
	}
	response := Response{
		Data:       data,
		Message:    "",
		Errors:     "",
		ErrorCode:  "",
		StatusCode: http.StatusOK,
		Success:    true,
		TraceID:    c.Keys[constant.TraceID].(string),
		MetaData:   paginationMetaData,
	}
	c.JSON(http.StatusOK, response)
}
