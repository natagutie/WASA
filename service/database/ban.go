package database

func (db *appdbimpl) SetBan(ban Ban) error {
	_, err := db.c.Exec("INSERT INTO ban (userID, banUserID) VALUES (?, ?)", ban.UserID, ban.BanUserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) SetUnBan(unban Ban) error {
	_, err := db.c.Exec("DELETE FROM ban WHERE banUserID=? AND userID=? ", unban.BanUserID, unban.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) BanCheck(userID1 Ban) (bool, error) {
	var isBan bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM ban WHERE userID=? AND banUserID=?)", userID1.UserID, userID1.BanUserID).Scan(&isBan)
	if err != nil {
		return false, err
	}
	return isBan, nil
}

func (db *appdbimpl) GetBanCount(userID string) (int, error) {
	var banCount int
	err := db.c.QueryRow("SELECT COUNT(*) FROM ban WHERE userID = ?", userID).Scan(&banCount)

	if err != nil {
		return 0, err
	}

	return banCount, nil
}
func (db *appdbimpl) GetBans(userID string) ([]string, error) {
	var bans []string
	rows, err := db.c.Query("SELECT banUserID FROM ban WHERE userID= ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var banUserID string
		err = rows.Scan(&banUserID)
		if err != nil {
			return nil, err
		}
		username, err := db.GetUsernameWithUserID(banUserID)
		if err != nil {
			return nil, err
		}

		bans = append(bans, username)
	}

	return bans, nil
}
