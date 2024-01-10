package api

import "strings"

func getAuth(bearer string) string {
	split := strings.Split(bearer, " ")
	userID := split[1]

	return userID
}
