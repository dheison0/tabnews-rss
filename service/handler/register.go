package handler

import (
    "github.com/gin-gonic/gin"
    "tabrss/service/database"
)

type ServiceHandler func(*gin.Context, database.Database)

func addDatabaseParameter(handler ServiceHandler, db database.Database) gin.HandlerFunc {
    return func (c *gin.Context) {
        handler(c, db)
    }
}

func AddHandlers(s *gin.Engine, db database.Database) {
    s.GET("/", addDatabaseParameter(GetUsers, db))
    s.POST("/", addDatabaseParameter(AddUser, db))
}
