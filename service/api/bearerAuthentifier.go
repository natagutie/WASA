package api

import "strings"

func getUserToken(authentication string) string {
	// We split the "Bearer" with the userID.
	userID := strings.Split(authentication, " ")
	// Take the userID part
	token_without_bearer := userID[1]

	return token_without_bearer
}