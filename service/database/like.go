package database

func (db *appdbimpl) LikePhoto(like Likes) error {
	_, err := db.c.Exec("INSERT INTO likes (userID, photoID) VALUES (?, ?)", like.UserID, like.PhotoID)
	if err != nil {
		return err
	}

	return  nil
}

func (db *appdbimpl) UnlikePhoto(like Likes) error {
	_, err := db.c.Exec("DELETE FROM likes WHERE photoID=? AND userID=? ", like.PhotoID, like.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) IfLike(like Likes) (bool, error) {
	var liker bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE photoID=? AND userID=?)", like.PhotoID, like.UserID).Scan(&liker)
	if err != nil {
		return false, err
	}
	return liker, nil

}