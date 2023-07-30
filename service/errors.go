package service

import "errors"

var (
	ErrUserNotFound = errors.New("this user doesn't exists anymore")
)
