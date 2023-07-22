package service

import (
	"fmt"
	"net/http"
)

func TabnewsUserExists(user string) (bool, error) {
	response, err := http.Get(fmt.Sprintf("%s/contents/%s?per_page=1", API_BASE, user))
	if err != nil {
		return false, err
	}
	return response.StatusCode==http.StatusOK, nil
}
