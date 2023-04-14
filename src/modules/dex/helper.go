package dex

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type IDexHelper interface {
	TestHelper()
}

type Helper struct {
	c *gin.Context
}

func NewDexHelper(c *gin.Context) *Helper {
	return &Helper{c: c}
}

func (dexHelper *Helper) TestHelper() {
	fmt.Println("TestHelper")
}
