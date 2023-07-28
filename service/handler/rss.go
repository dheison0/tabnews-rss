package handler

import (
	"net/http"
	"tabrss/service/database"
	"tabrss/service/feed"

	"github.com/gin-gonic/gin"
)

func FeedHandler(c *gin.Context, db database.Database) {
	feed := feed.GenerateFeed(db)
	rss, err := feed.ToRss()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to parse feed to RSS format"},
		)
		return
	}
	c.Data(http.StatusOK, "application/xml", []byte(rss))
}
