package feed

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"tabrss/service"
	"tabrss/service/database"
	"time"
)

type Post struct {
	Title     string    `json:"title"`
	Owner     string    `json:"owner_username"`
	CreatedAt time.Time `json:"created_at"`
	Slug      string    `json:"slug"`
}

func GetUserPosts(user string) ([]Post, error) {
	response, err := http.Get(service.API_BASE + "/contents/" + user)
	if err != nil {
		return []Post{}, err
	}
	defer response.Body.Close()
	var posts []Post
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return []Post{}, err
	}
	if err = json.Unmarshal(data, &posts); err != nil {
		return []Post{}, err
	}
	return posts, nil
}

func GetPostsFromUsers(users []database.User) [][]Post {
	allPosts := make([][]Post, 0, len(users))
	postsReceiver := make(chan []Post)
	for i := range users {
		go func(userIndex int, user string) {
			posts, err := GetUserPosts(user)
			if err != nil {
				log.Printf("Failed to get user '%s' posts: %s\n", users[userIndex].Name, err.Error())
				postsReceiver <- []Post{}
				return
			}
			postsReceiver <- posts
		}(i, users[i].Name)
	}
	for len(allPosts) != len(users) {
		allPosts = append(allPosts, <-postsReceiver)
	}
	return allPosts
}
