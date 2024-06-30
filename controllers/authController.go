package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login() func(g *gin.Context) {

	return func(g *gin.Context) {
		fmt.Println(g.Params)
	}
}
