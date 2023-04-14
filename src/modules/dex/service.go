package dex

import (
	"go-server/src/configs"
	"go-server/src/models/repositories"
	"go-server/src/modules/dex/dto"
	"go-server/src/share/base/base_model/pagination"
	basevalidator "go-server/src/share/base/base_validator"
	getLogger "go-server/src/share/logger"
	response "go-server/src/share/response"
	"go-server/src/share/utils"

	"github.com/gin-gonic/gin"
)

func GetDexes(c *gin.Context) {
	logger := getLogger.GetLogger().Logging
	logger.Info(utils.TID(c), "Logger")
	// Validator
	var listDexDto dto.ListDexDto
	// validator query
	basevalidator.ValidatorQuery(c, &listDexDto)
	// validator pagination query
	pagination.ParsePaginationDto(c, &listDexDto.PaginationDto)

	var dexRepository = repositories.NewDexRepository(configs.GetDB())
	var helper IDexHelper = NewDexHelper(c)
	helper.TestHelper()
	if listDexDto.Channel != "" {
		dexes, totalDex, err := dexRepository.ListDexesWithPaginationChannel(listDexDto)
		if err != nil {
			panic(err)
		}
		serializer := &DexesSerializer{c, dexes, dexRepository}
		response.ReturnDataWithPagination(c, serializer.Response(), totalDex)
		return
	}
	dexes, totalDex, err := dexRepository.ListDexesWithPagination(listDexDto)
	if err != nil {
		panic(err)
	}
	serializer := &DexesSerializer{c, dexes, dexRepository}
	response.ReturnDataWithPagination(c, serializer.Response(), totalDex)
}
