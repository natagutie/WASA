package database

func (db *appdbimpl) InsertComment(com Comments) (Comments, error) {
	_, err := db.c.Exec("INSERT INTO comments (userID,  username, photoID, comment) VALUES (?, ?, ?, ?)", com.UserID, com.Username, com.PhotoID, com.Comment)
	if err != nil {
		return Comments{}, err
	}
	return com, nil
}

func (db *appdbimpl) DeleteComment(com Comments) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE commentID=?", com.CommentID)
	if err != nil {
		return err
	}
	return nil
}


func (db *appdbimpl) GetPhotoComments(photoID uint64) ([]Comments, error) {
	var commentList []Comments
	rows, err := db.c.Query("SELECT userID, username, commentID, photoID, comment FROM comments WHERE photoID = ?", photoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// Commentol is each comment instance from Comment list
		var commentol Comments
		err = rows.Scan(&commentol.UserID, &commentol.Username, &commentol.CommentID, &commentol.PhotoID, &commentol.Comment)
		if err != nil {
			return nil, err
		}

		commentList = append(commentList, commentol)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return commentList, nil
}