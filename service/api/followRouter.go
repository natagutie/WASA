package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getAuth(r.Header.Get("Authorization"))

	followUsername := ps.ByName("followUsername")
	followUserID, err := rt.db.GetUserIDWithUsername(followUsername)
	if err != nil {
		return
	}
	var follow database.Follow
	follow.UserID = userID
	follow.FollowedID = followUserID

	isFollow, err := rt.db.IsFollowing(follow)
	if err != nil {
		return
	}
	if !isFollow {
		_ = rt.db.SetFollow(follow)
		w.WriteHeader(http.StatusCreated)

	}

}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	getAuth(r.Header.Get("Authorization"))

	followUsername := ps.ByName("followUsername")
	followUserID, err := rt.db.GetUserIDWithUsername(followUsername)
	if err != nil {
		http.Error(w, "cant get userID with username", http.StatusBadRequest)

		return
	}

	username := ps.ByName("username")
	userrID, err := rt.db.GetUserIDWithUsername(username)
	if err != nil {
		http.Error(w, "cant get userID with username", http.StatusBadRequest)

		return
	}
	var follow database.Follow
	follow.UserID = userrID
	follow.FollowedID = followUserID

	isFollow, err := rt.db.IsFollowing(follow)
	if err != nil {
		return
	}

	if isFollow {
		_ = rt.db.RemoveFollow(follow)

		message := "User unfollowed successfully"
		w.Write([]byte(message))
	}

}

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getAuth(r.Header.Get("Authorization"))

	followers, err := rt.db.GetFollowers(userID)
	if err != nil {
		http.Error(w, "cant get list of followers", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(followers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) getFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getAuth(r.Header.Get("Authorization"))
	w.Header().Set("content-type", "application/json")

	followers, err := rt.db.GetFollowings(userID)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(followers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) isFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getAuth(r.Header.Get("Authorization"))

	w.Header().Set("content-type", "application/json")

	followUsername := ps.ByName("followUsername")
	followUserID, err := rt.db.GetUserIDWithUsername(followUsername)
	if err != nil {
		http.Error(w, "cant find userID with username", http.StatusBadRequest)

		return
	}
	username := ps.ByName("username")
	userrID, err := rt.db.GetUserIDWithUsername(username)
	if err != nil {
		http.Error(w, "cant find userID with username", http.StatusBadRequest)

		return
	}
	var follow database.Follow
	if userrID == userID {
		follow.UserID = userID
	}
	follow.UserID = userrID

	follow.FollowedID = followUserID

	isFollow, err := rt.db.IsFollowing(follow)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(isFollow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
