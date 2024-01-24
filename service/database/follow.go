package database

func (db *appdbimpl) FollowUser(follow Follows) error {
	_, err := db.c.Exec("INSERT INTO follows (username, fUsername) VALUES (?, ?)", follow.Username, follow.FUsername)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) UnFollowUser(follow Follows) error {
	_, err := db.c.Exec("DELETE FROM follows WHERE fUsername=? AND username=?", follow.FUsername, follow.Username)
	if err != nil {
		return err
	}
	return nil
}



func (db *appdbimpl) IfFollow(follow Follows) (bool, error) {
	var follower bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM follows WHERE username=? AND fUsername=?)", follow.Username, follow.FUsername).Scan(&follower)
	if err != nil {
		return false, err
	}
	return follower, nil

}

func (db *appdbimpl) GetUserFollowers(username string) ([]string, error) {
	var followerList []string
	rows, err := db.c.Query("SELECT username FROM follows WHERE fUsername = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var usr string

		err = rows.Scan(&usr)
		if err != nil {
			return nil, err
		}
		followerList = append(followerList, usr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return followerList, nil
}

func (db *appdbimpl) GetUserFollowings(username string) ([]string, error) {
	var followingsList []string
	rows, err := db.c.Query("SELECT fUsername FROM follows WHERE username= ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var usr string
		err = rows.Scan(&usr)
		if err != nil {
			return nil, err
		}

		followingsList = append(followingsList, usr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return followingsList, nil
}