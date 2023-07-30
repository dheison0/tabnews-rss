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
		return []Post{}, service.ErrUserNotFound
	}
	return posts, nil
}

// GetPostsFromUsers retrieves posts from multiple users.
//
// Takes in a slice of User structs representing the users
// whose posts need to be retrieved. Returns a slice of slices
// of Post structs representing the posts of each user, and a
// slice of strings representing the usernames of users that
// were not found.
func GetPostsFromUsers(users []database.User) (allPosts [][]Post, usersNotFound []string) {
	allPosts = make([][]Post, 0, len(users))
	postsReceiver := make(chan []Post)
	for i := range users {
		userName := users[i].Name
		go func() {
			posts, err := GetUserPosts(userName)
			if err == service.ErrUserNotFound {
				usersNotFound = append(usersNotFound, userName)
			} else if err != nil {
				log.Printf("Failed to get posts from '%s': %s\n", userName, err.Error())
			}
			postsReceiver <- posts
		}()
	}
	for len(allPosts) != len(users) {
		allPosts = append(allPosts, <-postsReceiver)
	}
	return allPosts, usersNotFound
}
