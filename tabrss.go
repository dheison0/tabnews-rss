package main

import (
    "github.com/gin-gonic/gin"
    "tabrss/service/handler"
    "tabrss/service/database"
)

func main() {
    server := gin.Default()
    db := &database.SQLITEDatabase{}
    if err := db.Open("tablocal.db"); err != nil {
        panic(err)
    }
    defer db.Close()
    handler.AddHandlers(server, db)
    if err := server.Run(":8080"); err != nil {
        panic(err)
    }
}

