/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in webapi executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string conf:""
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	BanUser(ban Bans) error
	UnBanUser(ban Bans) error
	IfBan(ban Bans) (bool, error)
	GetBansList(username string) ([]string, error)

	InsertComment(com Comments) (Comments, error)
	DeleteComment(com Comments) error
	GetPhotoComments(photoID uint64) ([]Comments, error)

	FollowUser(follow Follows) error
	UnFollowUser(follow Follows) error
	IfFollow(follow Follows) (bool, error)
	GetUserFollowers(username string) ([]string, error)
	GetUserFollowings(username string) ([]string, error)

	LikePhoto(like Likes) (Likes, error)
	UnlikePhoto(like Likes) error
	IfLike(like Likes) (bool, error)

	UploadPhoto(picture Photos) error
	DeletePhoto(p Photos) error
	GetPhotoCount(userID string) (int, error)
	GetPhotos(picture Photos) ([]Photos, error)

	GetStream(username string) ([]Stream, error)
	LogUser(user Users) (string, error)
	SetNewUsername(username string, userID string) (string, error)

	GetUserID(username string) (string, error)
	UserExist(username string) (bool, error)
	GetUsername(userID string) (string, error)
	UserFollowingsCount(username string) (int, error)
	UserFollowersCount(username string) (int, error)
	BanCounts(username string) (int, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection db.
// db is required - an error will be returned if db is nil.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(SELECT name FROM sqlite_master WHERE type='table' AND name='users';).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		userData := `CREATE TABLE IF NOT EXISTS users (
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			userID TEXT NOT NULL UNIQUE,
			username TEXT NOT NULL
			);`
		_, err = db.Exec(userData)
		if err != nil {
			return nil, fmt.Errorf("error creating user database structure: %w", err)
		}
		followData := `CREATE TABLE follows (
			followID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			fUsername TEXT NOT NULL UNIQUE
			);`
		_, err = db.Exec(followData)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		photoData := `CREATE TABLE photos (
			photoID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			userID TEXT NOT NULL ,
			username TEXT NOT NULL ,
			date	TEXT NOT NULL,
			photo	BLOB,
			FOREIGN KEY (userID) REFERENCES users(userID)
			);`

		_, err = db.Exec(photoData)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		banData := `CREATE TABLE bans (
			banID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			bUsername TEXT NOT NULL
			);`
		_, err = db.Exec(banData)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		commentData := `CREATE TABLE comments (
			commentID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			userID TEXT NOT NULL,
			username TEXT NOT NULL,
			photoID INTEGER NOT NULL,
			comment	TEXT NOT NULL,
			FOREIGN KEY (userID) REFERENCES users(userID),
			FOREIGN KEY (photoID) REFERENCES photos(photoID)
			);`
		_, err = db.Exec(commentData)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		likeData := `CREATE TABLE likes (
			likeID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			userID TEXT NOT NULL,
			photoID INTEGER NOT NULL,
			FOREIGN KEY (userID) REFERENCES users(userID),
			FOREIGN KEY (photoID) REFERENCES photos(photoID)
			);`
		_, err = db.Exec(likeData)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}