package handler

import (
	"log"
	"net/http"
	"tabrss/service"
	"tabrss/service/database"

	"github.com/gin-gonic/gin"
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
	userExists, err := service.TabnewsUserExists(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gin.H{"error": "failed to check if user exists"})
		return
	} else if !userExists {
		c.JSON(http.StatusBadRequest, &gin.H{"error": "user not exists"})
		return
	}
	userWasAdded := db.AddUser(database.User{Name: user, Exists: true})
	if userWasAdded {
		c.JSON(http.StatusOK, &gin.H{"message": "user added"})
		return
	}
	c.JSON(http.StatusBadRequest, &gin.H{"error": "user already added"})
}

func GetUsers(c *gin.Context, db database.Database) {
	users, err := db.GetUsers()
	if err != nil {
		log.Printf("Failed to get users: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, &gin.H{"error": "failed to get users"})
		return
	}
	c.JSON(http.StatusOK, &gin.H{"message": "ok", "users": &users})
}

func RemoveUser(c *gin.Context, db database.Database) {
	user := c.Query("user")
	if user == "" {
		c.JSON(http.StatusBadRequest, &gin.H{"error": "query 'user' is missing"})
		return
	}
	err := db.RemoveUser(database.User{Name: user})
	if err == nil {
		c.JSON(http.StatusOK, &gin.H{"message": "user removed"})
	} else {
		c.JSON(http.StatusInternalServerError, &gin.H{"error": "failed to remove user"})
	}
}
