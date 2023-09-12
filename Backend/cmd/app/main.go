package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Game-as-a-Service/The-Message/app/delivery/game"
)

func main() {
	router := gin.Default()
	router.GET("/gameInit", game.GameInit)

	router.Run("0.0.0.0:8080")
}
