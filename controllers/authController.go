package controllers

import (
	"github.com/gin-gonic/gin"
)

func Login() func(g *gin.Context) {

	return func(g *gin.Context) {

		g.JSON(200, gin.H{
			"test": "working",
		})

	}
}
