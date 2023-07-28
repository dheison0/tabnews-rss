package feed

import (
	"fmt"
	"tabrss/service"
	"tabrss/service/database"
	"time"

	"github.com/gorilla/feeds"
)

func GenerateFeed(db database.Database) feeds.Feed {
	feed := feeds.Feed{
		Title:       service.FEED_TITLE,
		Description: service.FEED_DESCRIPTION,
		Link:        &feeds.Link{Href: service.FEED_LINK},
		Created:     time.Now(),
	}
	users, err := db.GetUsers()
	if err != nil {
		return feed
	}
	var items []*feeds.Item
	posts := GetPostsFromUsers(users)
	for byUser := range posts {
		for i := range posts[byUser] {
			post := posts[byUser][i]
			if post.Title == "" {
				continue
			}
			items = append(items, TurnPostIntoFeedItem(post))
		}
	}
	feed.Items = items
	return feed
}

func TurnPostIntoFeedItem(post Post) *feeds.Item {
	return &feeds.Item{
		Title:   post.Title,
		Created: post.CreatedAt,
		Link: &feeds.Link{
			Href: fmt.Sprintf("%s/%s/%s", service.SITE, post.Owner, post.Slug),
		},
	}
}
