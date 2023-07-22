package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tabrss/service"
	"tabrss/service/database"
)

func AddUser(c *gin.Context, db database.Database) {
	user := c.Query("user")
	if user == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "query 'user' is missing"},
		)
		return
	}
	exists, err := service.TabnewsUserExists(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check if user exists"})
		return
	} else if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not exists"})
		return
	}
	if ok := db.AddUser(database.User{Name: user, Exists: true}); ok {
		c.JSON(http.StatusOK, gin.H{"message": "user added"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "user already added"})
}

func GetUsers(c *gin.Context, db database.Database) {
	users, err := db.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "users": users})
}
