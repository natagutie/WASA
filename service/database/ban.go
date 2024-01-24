package database

func (db *appdbimpl) BanUser(ban Bans) error {
	_, err := db.c.Exec("INSERT INTO bans (username, bUsername) VALUES (?, ?)", ban.Username, ban.BUsername)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) UnBanUser(ban Bans) error {
	_, err := db.c.Exec("DELETE FROM bans WHERE username=? AND bUsername=? ", ban.Username, ban.BUsername)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) IfBan(ban Bans) (bool, error) {
	var ifBan bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM bans WHERE username=? AND bUsername=?)", ban.Username, ban.BUsername).Scan(&ifBan)
	if err != nil {
		return false, err
	}
	return ifBan, nil
}

func (db *appdbimpl) GetBansList(username string) ([]string, error) {
	var banList []string
	rows, err := db.c.Query("SELECT bUsername FROM bans WHERE username= ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var bUsername string
		err = rows.Scan(&bUsername)
		if err != nil {
			return nil, err
		}
		banList = append(banList, username)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return banList, nil
}