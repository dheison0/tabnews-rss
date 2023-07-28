package main

import (
	"tabrss/service/database"
	"tabrss/service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db := &database.SQLITEDatabase{}
	if err := db.Open("tablocal.db"); err != nil {
		panic(err)
	}
	defer db.Close()
	handler.AddHandlers(server, db)
	server.StaticFile("/", "web/index.html")
	server.Static("/web", "web")
	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}
