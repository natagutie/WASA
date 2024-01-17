package database

// set pic
func (db *appdbimpl) SetPic(pic Photo) error {
	_, err := db.c.Exec("INSERT INTO photos (userID, username, date, photo) VALUES (?, ?, ?, ?)", pic.UserID, pic.Username, pic.Date, pic.Photo)
	if err != nil {
		return err // Return if error
	}
	return nil // Void for no error
}

// Remove pic
func (db *appdbimpl) RemovePic(picID Photo) error {
	// Delete from the photo table
	_, err := db.c.Exec("DELETE FROM photos WHERE photoID=?", picID.PhotoID)
	if err != nil {
		return err // Return error
	}

	// Delete from the like table
	_, err = db.c.Exec("DELETE FROM like WHERE photoID=?", picID.PhotoID)
	if err != nil {
		return err
	}

	// Delete from the comment table
	_, err = db.c.Exec("DELETE FROM comment WHERE photoID=?", picID.PhotoID)
	if err != nil {
		return err
	}

	return nil // void if no errors

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
