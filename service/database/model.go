package database

type User struct {
	Name   string `json:"name"`
	Exists bool   `json:"exists"`
}

type Database interface {
	Open(string) error
	Close()
	GetUsers() ([]User, error)
	AddUser(User) bool
	RemoveUser(User) error
}
