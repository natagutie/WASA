--
-- SQL file to initialize the database.
-- Database schema version: 0
--

CREATE TABLE IF NOT EXISTS User
(
	id       BLOB NOT NULL PRIMARY KEY,
	name     TEXT NOT NULL,
	surname  TEXT NOT NULL,
	username TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Photo
(
	id          BLOB NOT NULL PRIMARY KEY,
	imageUrl    TEXT NOT NULL,
	authorId    BLOB NOT NULL REFERENCES User (id),
	publishDate TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Comment
(
	id          BLOB NOT NULL PRIMARY KEY,
	`text`      TEXT NOT NULL,
	publishDate TEXT NOT NULL,
	authorId    BLOB NOT NULL REFERENCES User (id) ON DELETE CASCADE,
	photoId     BLOB NOT NULL REFERENCES Photo (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Likes
(
	userId  BLOB NOT NULL REFERENCES User (id) ON DELETE CASCADE,
	photoId BLOB NOT NULL REFERENCES Photo (id) ON DELETE CASCADE,
	PRIMARY KEY (userId, photoId)
);

CREATE TABLE IF NOT EXISTS Follow
(
	followerId BLOB NOT NULL REFERENCES User (id) ON DELETE CASCADE,
	followedId BLOB NOT NULL REFERENCES User (id) ON DELETE CASCADE,
	PRIMARY KEY (followerId, followedId)
);

CREATE TABLE IF NOT EXISTS Ban
(
	bannerId BLOB NOT NULL REFERENCES User (id) ON DELETE CASCADE,
	bannedId BLOB NOT NULL REFERENCES User (id) ON DELETE CASCADE,
	PRIMARY KEY (bannerId, bannedId)
);
