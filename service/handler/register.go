package handler

import (
	"tabrss/service/database"

	"github.com/gin-gonic/gin"
)

type ServiceHandler func(*gin.Context, database.Database)

func addDatabaseParameter(handler ServiceHandler, db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, db)
	}
}

func AddHandlers(s *gin.Engine, db database.Database) {
	s.GET("/api/user", addDatabaseParameter(GetUsers, db))
	s.POST("/api/user", addDatabaseParameter(AddUser, db))
	s.GET("/rss", addDatabaseParameter(FeedHandler, db))
}
