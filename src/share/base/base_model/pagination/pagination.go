package pagination

import (
	baseDto "server-go/src/share/base/base_dto"
	validator "server-go/src/share/base/base_validator"
	"server-go/src/share/constant"
	"server-go/src/share/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPaginateScope(pageSize, page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetPaginateParams(pageSize, page int) baseDto.PaginationParams {
	offset := (page - 1) * pageSize
	return baseDto.PaginationParams{
		Limit:  pageSize,
		Offset: offset,
	}
}

func ParsePaginationDto(c *gin.Context, dto *baseDto.PaginationDto) baseDto.PaginationParams {
	// validation
	validator.ValidatorQuery(c, dto)
	defaultLimit := 10
	defaultPage := 1
	dto.PageSize = utils.CoalesceInt(dto.PageSize, defaultLimit)
	dto.Page = utils.CoalesceInt(dto.Page, defaultPage)
	params := GetPaginateParams(dto.PageSize, dto.Page)
	dto.PaginationParams = params
	c.Set(constant.Pagination, *dto)
	return params
}
