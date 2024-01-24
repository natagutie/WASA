package database

import (
	"github.com/gofrs/uuid"
)

func (db *appdbimpl) LogUser(user Users) (string, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	user.UserID = id.String()

	_, err = db.c.Exec("INSERT INTO users(username, userID) VALUES (?, ?)", user.Username, user.UserID)
	if err != nil {
		return "", err
	}
	return user.UserID, nil

}

func (db *appdbimpl) SetNewUsername(username string, userID string) (string, error) {
	_, err := db.c.Exec("UPDATE users SET username=? WHERE userID=?", username, userID)
	if err != nil {
		return "", err
	}
	return username, nil
}