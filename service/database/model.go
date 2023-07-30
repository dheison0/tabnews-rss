package database

type User struct {
	Name   string `json:"name"`
	Exists bool   `json:"exists"`
}

type Database interface {
	AddUser(User) bool
	Close()
	GetUsers() ([]User, error)
	Open(string) error
	RemoveUser(User) error
	SetUserExists(User, bool) error
}
