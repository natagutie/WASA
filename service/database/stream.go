package database

func (db *appdbimpl) GetStream(username string, userID string) ([]Stream, error) {

	rows, err := db.c.Query("SELECT p.photoID, p.userID, p.photo, p.username, p.date FROM photos p INNER JOIN follows f ON p.username = f.fUsername WHERE f.username = ? ORDER BY p.date DESC", username)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var streamList []Stream

	for rows.Next() {
		var stream Stream
		err := rows.Scan(&stream.PhotoID, &stream.UserID, &stream.Photo, &stream.FUsername, &stream.Date)
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow("SELECT COUNT(*) FROM likes WHERE photoID=?", stream.PhotoID).Scan(&stream.LikeCount)
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow("SELECT COUNT(*) FROM comments WHERE photoID=?", stream.PhotoID).Scan(&stream.CommentCount)
		if err != nil {
			return nil, err
		}

		err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE userID=? AND photoID=?)", userID, stream.PhotoID).Scan(&stream.IfLike)
		if err != nil {
			return nil, err
		}

		commentList, err := db.GetPhotoComments(stream.PhotoID)
		if err != nil {
			return nil, err
		}
		stream.Comments = commentList

		streamList = append(streamList, stream)

	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return streamList, nil
}