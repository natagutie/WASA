package database

func (db *appdbimpl) SetComment(c Comment) (Comment, error) {
	_, err := db.c.Exec("INSERT INTO comment (userID, username, photoID, comment) VALUES (?, ?, ?, ?)", c.UserID, c.Username, c.PhotoID, c.Comment)
	if err != nil {
		return Comment{}, err
	}
	return c, nil
}

func (db *appdbimpl) RemoveComment(c Comment) error {
	_, err := db.c.Exec("DELETE FROM comment WHERE commentID=?", c.CommentID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveBanComment(ban Ban) error {
	_, err := db.c.Exec("DELETE FROM comment WHERE userID=?", ban.BanUserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetCommentCount(pic Photo) (int, error) {
	var countNumb int
	err := db.c.QueryRow("SELECT COUNT(*) FROM comment WHERE photoID=?", pic.PhotoID).Scan(&countNumb)
	if err != nil {
		return 0, err
	}
	return countNumb, nil
}

func (db *appdbimpl) GetPhotoUsernameWithPhotoID(pic Photo) (int, error) {
	var countNumb int
	err := db.c.QueryRow("SELECT username FROM comment WHERE photoID=?", pic.PhotoID).Scan(&countNumb)
	if err != nil {
		return 0, err
	}
	return countNumb, nil
}
func (db *appdbimpl) GetComments(pic uint64) ([]Comment, error) {
	var comment []Comment
	rows, err := db.c.Query("SELECT commentID, userID, comment FROM comment WHERE photoID = ?", pic)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var commenter Comment
		err = rows.Scan(&commenter.CommentID, &commenter.UserID, &commenter.Comment)
		if err != nil {
			return nil, err
		}

		Username, err := db.GetUsernameWithUserID(commenter.UserID)
		if err != nil {
			return nil, err
		}
		commenter.Username = Username

		comment = append(comment, commenter)
	}

	return comment, nil
}
