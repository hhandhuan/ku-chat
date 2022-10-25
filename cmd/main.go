package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"ku-chat/internal/router"
	"ku-chat/pkg/config"
	"log"
)

func main() {
	engine := gin.Default()
	engine.Static("/assets", "../assets")
	engine.LoadHTMLGlob("../view/*")

	store := cookie.NewStore([]byte(config.Conf.Session.Secret))
	engine.Use(sessions.Sessions(config.Conf.Session.Name, store))

	router.RegisterWebRouter(engine)
	router.RegisterWsRouter(engine)

	if err := engine.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
