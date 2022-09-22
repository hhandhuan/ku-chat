package main

import (
	"github.com/gin-gonic/gin"
	"ku-chat/internal/router"
	"log"
)

func main() {
	engine := gin.Default()
	engine.Static("/assets", "../assets")
	engine.LoadHTMLGlob("../view/*")
	router.RegisterRouter(engine)
	router.RegisterWsRouter(engine)
	if err := engine.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
