package database

func (db *appdbimpl) SetLike(like Like) (Like, error) {
	liker, err := db.c.Exec("INSERT INTO like (userID, photoID) VALUES (?, ?)", like.UserID, like.PhotoID)
	if err != nil {
		return Like{}, err
	}

	likeID, err := liker.LastInsertId()
	if err != nil {
		return Like{}, err
	}
	like.LikeID = likeID
	return like, nil
}

func (db *appdbimpl) RemoveLike(like Like) error {
	_, err := db.c.Exec("DELETE FROM like WHERE userID=? AND photoID=?", like.UserID, like.PhotoID)
	if err != nil {
		return err
	}
	return nil
}
func (db *appdbimpl) GetLikeCount(pic Photo) (int, error) {
	var likeNumb int
	err := db.c.QueryRow("SELECT COUNT(*) FROM like WHERE photoID=?", pic.PhotoID).Scan(&likeNumb)
	if err != nil {
		return 0, err
	}
	return likeNumb, nil
}

func (db *appdbimpl) IsLiked(like Like) (bool, error) {
	var isliked bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM like WHERE userID=? AND photoID=?)", like.UserID, like.PhotoID).Scan(&isliked)
	if err != nil {
		return false, err
	}
	return isliked, nil

}

func (db *appdbimpl) GetPhotoCount(userID string) (int, error) {
	var photoCount int
	err := db.c.QueryRow("SELECT COUNT(*) FROM photos WHERE userID = ?", userID).Scan(&photoCount)

	if err != nil {
		return 0, err
	}

	return photoCount, nil
}

func (db *appdbimpl) GetPhotos(pixel Photo) ([]Photo, error) {
	var photo []Photo
	rows, err := db.c.Query("SELECT photoID, userID, date, photo FROM photos WHERE userID = ? ORDER BY date DESC", pixel.PhotoUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&pixel.PhotoID, &pixel.PhotoUserID, &pixel.Date, &pixel.Photo)
		if err != nil {
			return nil, err
		}
		likeCount, err := db.GetLikeCount(pixel)
		if err != nil {
			return nil, err
		}
		pixel.LikesNum = likeCount

		commentCount, err := db.GetCommentCount(pixel)
		if err != nil {
			return nil, err
		}
		pixel.CommentNum = commentCount

		var like Like
		like.PhotoID = pixel.PhotoID
		like.UserID = pixel.UserID
		isliked, err := db.IsLiked(like)
		if err != nil {
			return nil, err
		}
		pixel.IsLiked = isliked

		comments, err := db.GetComments(pixel.PhotoID)
		if err != nil {
			return nil, err
		}
		pixel.Comments = comments

		photo = append(photo, pixel)

	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return photo, nil

}



func (db *appdbimpl) RemovePhoto( photo Photo) error {
	// Delete from the photo table
	_, err := db.c.Exec("DELETE FROM photos WHERE photoID=?", photo.PhotoID)
	if err != nil {
		return err // Return error
	}

	// Delete from the like table
	_, err = db.c.Exec("DELETE FROM like WHERE photoID=?", photo.PhotoID)
	if err != nil {
		return err
	}

	// Delete from the comment table
	_, err = db.c.Exec("DELETE FROM comment WHERE photoID=?", photo.PhotoID)
	if err != nil {
		return err
	}

	return nil // void if no errors

}
func (db *appdbimpl) InsertPhoto(pic Photo) error {
	_, err := db.c.Exec("INSERT INTO photos (userID, username, date, photo) VALUES (?, ?, ?, ?)", pic.UserID, pic.Username, pic.Date, pic.Photo)
	if err != nil {
		return err // Return if error
	}
	return nil // Void for no error
}