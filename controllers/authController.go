package controllers

import (
	"github.com/gin-gonic/gin"
)

func Login() func(g *gin.Context) {

	return func(g *gin.Context) {
		g.Header("Access-Control-Allow-Origin", "*")
		g.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")

		g.JSON(200, gin.H{
			"test": "working",
		})

	}
}
