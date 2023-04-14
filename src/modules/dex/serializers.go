package dex

import (
	"go-server/src/models/entities"
	"go-server/src/models/repositories"

	"github.com/gin-gonic/gin"
)

type Serializer struct {
	c          *gin.Context
	dex        entities.DexEntity
	repository *repositories.DexRepository
}

type Response = entities.DexEntity

func (serializer *Serializer) Response() Response {
	// handle here
	res := serializer.dex
	return res
}

type DexesSerializer struct {
	c          *gin.Context
	Dexes      []entities.DexEntity
	repository *repositories.DexRepository
}

func (dexesSerializer *DexesSerializer) Response() []Response {
	var response []Response
	// nolint
	for _, dex := range dexesSerializer.Dexes {
		serializer := Serializer{dexesSerializer.c, dex, dexesSerializer.repository}
		response = append(response, serializer.Response())
	}
	return response
}
