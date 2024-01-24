package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)



func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))
	var unfollow database.Follows

	fUsername := ps.ByName("fUsername")
	username, err := rt.db.GetUsername(userID)
	if err != nil {
		http.Error(w, "cant get userID with username", http.StatusBadRequest)
		return
	}

	unfollow.UserID = userID
	unfollow.Username = username
	unfollow.FUsername = fUsername

	err = rt.db.UnFollowUser(unfollow)
	if err != nil {
		http.Error(w, "cannot unfollow :(", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
func (rt *_router) ifFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))
	username, err := rt.db.GetUsername(userID)
	if err != nil {
		http.Error(w, "cant find username", http.StatusBadRequest)

		return
	}
	fUsername := ps.ByName("fUsername")

	var follow database.Follows
	follow.UserID = userID
	follow.Username = username
	follow.FUsername = fUsername

	FollowChecker, err := rt.db.IfFollow(follow)
	if err != nil {
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(FollowChecker)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))

	fUsername := ps.ByName("fUsername")
	username, err := rt.db.GetUsername(userID)
	if err != nil {
		return
	}
	var follow database.Follows
	follow.UserID = userID
	follow.Username = username
	follow.FUsername = fUsername

	err = rt.db.FollowUser(follow)
	if err != nil {
		http.Error(w, "cant follow user", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (rt *_router) getUserFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))
	w.Header().Set("content-type", "application/json")

	username, err := rt.db.GetUsername(userID)
	if err != nil {
		http.Error(w, "cant find username", http.StatusBadRequest)

		return
	}

	followersList, err := rt.db.GetUserFollowers(username)
	if err != nil {
		http.Error(w, "cant get user followers list", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(followersList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) getUserFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))
	w.Header().Set("content-type", "application/json")

	username, err := rt.db.GetUsername(userID)
	if err != nil {
		http.Error(w, "cant find username", http.StatusBadRequest)

		return
	}
	followingList, err := rt.db.GetUserFollowings(username)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(followingList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}