package database

import (
	"log"

	"github.com/gofrs/uuid"
)


func (db *appdbimpl) LogUser(usr User) (string, error) {

	checkUser, err := db.CheckUserExist(usr.Username)

	if err != nil {
		return "error in checking user existing", nil
	}
	if checkUser {
		user, err := db.GetUserIDWithUsername(usr.Username)
		if err != nil {
			return "error retrieving userID with username", nil
		}
		return user, nil

	} else {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalf("couldnt generate userid : %v", err)
		}
		usr.UserID = id.String()

		_, err2 := db.c.Exec(INSERT INTO user(username, userID) VALUES (?, ?), usr.Username, usr.UserID)
		if err2 != nil {
			return "error!", err
		}
		return usr.UserID, nil
	}
}

func (db *appdbimpl) GetUsernameWithUserID(userID string) (string, error) {
	var username string

	err := db.c.QueryRow("SELECT username FROM user WHERE userID=?", userID).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}

func (db *appdbimpl) GetUserIDWithUsername(username string) (string, error) {
	var usrID string

	err := db.c.QueryRow("SELECT userID FROM user WHERE username=?", username).Scan(&usrID)
	if err != nil {
		return "did not work", err
	}

	return usrID, nil
}

func (db *appdbimpl) CheckUserExist(username string) (bool, error) {
	var userExists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 from user WHERE username=?)", username).Scan(&userExists)
	if err != nil {
		return false, err
	}
	return userExists, nil
}
func (db *appdbimpl) SetUsername(username string, user User) (string, error) {
	_, err := db.c.Exec("UPDATE user SET username=? WHERE  userID=?", username, user.UserID)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (db *appdbimpl) GetStream(user User) ([]Stream, error) {
	//select photos from followed users with likes and comments count, in reverse chronological order
	rows, err := db.c.Query("SELECT p.photoID, p.photo, p.userID, p.date FROM photos p INNER JOIN follow f ON p.userID = f.toFollowID WHERE f.userID = ? ORDER BY p.date DESC", user.UserID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var streams []Stream
	for rows.Next() {
		var s Stream
		err := rows.Scan(&s.PhotoID, &s.Photo, &s.FollowedUserID, &s.Date)
		if err != nil {
			return nil, err
		}
		followed_username, err := db.GetUsernameWithUserID(s.FollowedUserID)
		if err != nil {
			return nil, err
		}
		s.FollowedUsername = followed_username

		var pixel Photo
		pixel.PhotoID = s.PhotoID
		pixel.UserID = user.UserID
		likeCount, err := db.GetLikeCount(pixel)
		if err != nil {
			return nil, err
		}
		s.LikeCount = likeCount

		commentCount, err := db.GetCommentCount(pixel)
		if err != nil {
			return nil, err
		}
		s.CommentCount = commentCount
		s.UserID = user.UserID

		var like Like
		like.PhotoID = pixel.PhotoID
		like.UserID = pixel.UserID
		isliked, err := db.IsLiked(like)
		if err != nil {
			return nil, err
		}
		s.IsLiked = isliked

		comments, err := db.GetComments(pixel.PhotoID)
		if err != nil {
			return nil, err
		}
		s.Comments = comments

		streams = append(streams, s)
	}

	return streams, nil
}