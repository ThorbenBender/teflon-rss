package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	tokenString := header.Get("Authorization")
	if tokenString == "" {
		return "", errors.New("No api key found")
	}
	tokenArray := strings.Split(tokenString, " ")
	if tokenArray[0] != "ApiKey" || len(tokenArray) != 2 {
		return "", errors.New("Api Key in wrong format")
	}
	return tokenArray[1], nil
}
