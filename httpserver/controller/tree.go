package controller

import (
	"example-tools/httpserver/library/fp"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Tree(c *gin.Context) {
	fmt.Println("hello word !")
    c.JSON(200, "hello word !")

	fp.Fp("this is a test")
}

