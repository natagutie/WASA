package database

func (db *appdbimpl) GetUsername(userID string) (string, error) {
	var username string

	err := db.c.QueryRow("SELECT username FROM users WHERE userID=?", userID).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}
func (db *appdbimpl) GetUserID(username string) (string, error) {
	var userid string

	err := db.c.QueryRow("SELECT userID FROM users WHERE username=?", username).Scan(&userid)
	if err != nil {
		return "", err
	}

	return userid, nil
}
func (db *appdbimpl) UserExist(username string) (bool, error) {
	var exist bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 from users WHERE username=?)", username).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}
func (db *appdbimpl) UserFollowersCount(username string) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follows WHERE fUsername = ?", username).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *appdbimpl) UserFollowingsCount(username string) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follows WHERE username = ? ", username).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *appdbimpl) BanCounts(username string) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM bans WHERE username = ?", username).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}