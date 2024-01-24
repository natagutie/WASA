package database

func (db *appdbimpl) UploadPhoto(picture Photos) error {
	_, err := db.c.Exec("INSERT INTO photos (userID, username, date, photo) VALUES (?, ?, ?, ?)", picture.UserID, picture.PUsername, picture.Date, picture.Photo)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) DeletePhoto(p Photos) error {
	_, err := db.c.Exec("DELETE FROM photos WHERE photoID=?", p.PhotoID)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("DELETE FROM comments WHERE photoID=?", p.PhotoID)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("DELETE FROM likes WHERE photoID=?", p.PhotoID)
	if err != nil {
		return err
	}

	return nil

}

func (db *appdbimpl) GetPhotoCount(userID string) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM photos WHERE userID = ?", userID).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *appdbimpl) GetPhotos(picture Photos) ([]Photos, error) {
	var photo []Photos
	rows, err := db.c.Query("SELECT photoID, userID, date, photo FROM photos WHERE username = ? ORDER BY date DESC", picture.PUsername)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&picture.PhotoID, &picture.UserID, &picture.Date, &picture.Photo)
		if err != nil {
			return nil, err
		}
		err = db.c.QueryRow("SELECT COUNT(*) FROM likes WHERE photoID=?", picture.PhotoID).Scan(&picture.LikeCount)
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow("SELECT COUNT(*) FROM comments WHERE photoID=?", picture.PhotoID).Scan(&picture.CommentCount)
		if err != nil {
			return nil, err
		}
		err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE userID=? AND photoID=?)", picture.UserID, picture.PhotoID).Scan(&picture.IfLiked)
		if err != nil {
			return nil, err
		}

		comments, err := db.GetPhotoComments(picture.PhotoID)
		if err != nil {
			return nil, err
		}
		picture.Comments = comments

		photo = append(photo, picture)

	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return photo, nil

}