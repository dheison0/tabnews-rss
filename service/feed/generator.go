package feed

import (
	"fmt"
	"log"
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
		log.Printf("Failed to get users from database: %s\n", err.Error())
		return feed
	}
	var items []*feeds.Item
	posts, usersNotFound := GetPostsFromUsers(users)
	go SetUsersNotExistsAnymore(db, usersNotFound)
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
		Author:  &feeds.Author{Name: post.Owner},
		Created: post.CreatedAt,
		Link: &feeds.Link{
			Href: fmt.Sprintf("%s/%s/%s", service.SITE, post.Owner, post.Slug),
		},
	}
}

func SetUsersNotExistsAnymore(db database.Database, users []string) {
	for _, user := range users {
		if err := db.SetUserExists(database.User{Name: user}, false); err != nil {
			log.Printf("Failed to remove user '%s': %s\n", user, err.Error())
		}
	}
}
