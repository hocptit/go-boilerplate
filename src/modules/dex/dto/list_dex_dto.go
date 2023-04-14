package dto

import (
	baseDto "go-server/src/share/base/base_dto"
)

type ListDexDto struct {
	baseDto.PaginationDto
	SortBy    string `form:"sortBy"`
	Direction string `form:"direction"`
	Network   string `form:"network"`
	Channel   string `form:"channel"`
}
