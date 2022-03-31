package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	addRoutest(engine)
	log.Fatalln(engine.Run())
}

func addRoutest(engine *gin.Engine) {
	engine.GET("/health", GetHealth())
}

func GetHealth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
