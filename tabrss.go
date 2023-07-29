package main

import (
	"tabrss/service/database"
	"tabrss/service/handler"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	dbFile := os.Getenv("DATABASE_FILE")
	if port == "" {
		port = "8080"
	}
	if dbFile == "" {
		dbFile = "tablocal.db"
	}
	server := gin.Default()
	db := &database.SQLITEDatabase{}
	if err := db.Open(dbFile); err != nil {
		panic(err)
	}
	defer db.Close()
	handler.AddHandlers(server, db)
	server.StaticFile("/", "web/index.html")
	server.Static("/web", "web")
	if err := server.Run(":" + port); err != nil {
		panic(err)
	}
}
