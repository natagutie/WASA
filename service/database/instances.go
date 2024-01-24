package database

type Users struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
}

type Bans struct {
	Username  string `json:"username"`
	BUsername string `json:"bUsername"`
}

type Profile struct {
	UserID         string   `json:"userID"`
	Username       string   `json:"username"`
	Pictures       []Photos `json:"photoList"`
	PhotoCount     int      `json:"photoCount"`
	FollowingCount int      `json:"followingCount"`
	FollowCount    int      `json:"followerCount"`
	BanCount       int      `json:"banCount"`
	Bans           []string `json:"bansList"`
}

type Photos struct {
	PhotoID      uint64     `json:"photoID"`
	UserID       string     `json:"userID"`
	PUsername    string     `json:"pUsername"`
	Username     string     `json:"username"`
	Date         string     `json:"date"`
	LikeCount    int        `json:"likeCount"`
	CommentCount int        `json:"commentCount"`
	Comments     []Comments `json:"commentList"`
	Photo        []byte     `json:"photoFile"`
	IfLiked      bool       `json:"ifLiked"`
}

type Likes struct {
	UserID  string `json:"userID"`
	PhotoID uint64 `json:"photoID"`
	LikeID  int64  `json:"likeID"`
}

type Comments struct {
	CommentID     uint64 `json:"commentID"`
	UserID        string `json:"userID"`
	Username      string `json:"username"`
	PhotoUsername string `json:"pUsername"`
	PhotoUserID   string `json:"pUserID"`
	PhotoID       uint64 `json:"photoID"`
	Comment       string `json:"comment"`
}

type Stream struct {
	UserID       string     `json:"userID"`
	CommentCount int        `json:"commentCount"`
	Photo        []byte     `json:"photoFile"`
	Comments     []Comments `json:"commentList"`
	IfLike       bool       `json:"ifLike"`
	FUsername    string     `json:"fUsername"`
	Date         string     `json:"date"`
	PhotoID      uint64     `json:"photoID"`
	LikeCount    int        `json:"likeCount"`
}

type Follows struct {
	Username  string `json:"username"`
	FUsername string `json:"fUsername"`
	UserID    string `json:"userID"`
}
