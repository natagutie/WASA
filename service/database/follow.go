package database

func (db *appdbimpl) SetFollow(follow Follow) error {
	_, err := db.c.Exec("INSERT INTO follow ( userID, toFollowID) VALUES (?, ?)", follow.UserID, follow.FollowedID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveFollow(follow Follow) error {
	_, err := db.c.Exec("DELETE FROM follow WHERE toFollowID=? AND userID=?", follow.FollowedID, follow.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetFollowCount(userID string) (int, error) {
	var followCount int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follow WHERE toFollowID = ?", userID).Scan(&followCount)

	if err != nil {
		return 0, err
	}

	return followCount, nil
}

func (db *appdbimpl) GetFollowingCount(userID string) (int, error) {
	var followingCount int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follow WHERE userID = ? ", userID).Scan(&followingCount)

	if err != nil {
		return 0, err
	}

	return followingCount, nil
}

func (db *appdbimpl) IsFollowing(follow Follow) (bool, error) {
	var isFollowing bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM follow WHERE  userID=? AND toFollowID=?)", follow.UserID, follow.FollowedID).Scan(&isFollowing)
	if err != nil {
		return false, err
	}
	return isFollowing, nil

}

func (db *appdbimpl) GetFollowers(userID string) ([]string, error) {
	var followers []string
	rows, err := db.c.Query("SELECT userID FROM follow WHERE toFollowID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var followerUserID string
		err = rows.Scan(&followerUserID)
		if err != nil {
			return nil, err
		}

		username, err := db.GetUsernameWithUserID(followerUserID)
		if err != nil {
			return nil, err
		}
		followers = append(followers, username)
	}

	return followers, nil
}

func (db *appdbimpl) GetFollowings(userID string) ([]string, error) {
	var following []string
	rows, err := db.c.Query("SELECT toFollowID FROM follow WHERE userID= ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var followingUserID string
		err = rows.Scan(&followingUserID)
		if err != nil {
			return nil, err
		}
		username, err := db.GetUsernameWithUserID(followingUserID)
		if err != nil {
			return nil, err
		}

		following = append(following, username)
	}

	return following, nil
}
