package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	var profile database.Profile

	profile.Username = ps.ByName("pUsername")

	userID, err := rt.db.GetUserID(profile.Username)
	if err != nil {
		http.Error(w, "no userID", http.StatusBadRequest)
		return
	}
	profile.UserID = userID

	photocount, err := rt.db.GetPhotoCount(profile.UserID)
	if err != nil {
		http.Error(w, "error: photocount profile", http.StatusBadRequest)
		return
	}
	profile.PhotoCount = photocount

	bCount, err := rt.db.BanCounts(profile.Username)
	if err != nil {
		http.Error(w, "error getting ban count", http.StatusBadRequest)
		return
	}
	profile.BanCount = bCount

	bansList, err := rt.db.GetBansList(profile.Username)
	if err != nil {
		http.Error(w, "error getting bans", http.StatusBadRequest)
		return
	}
	profile.Bans = bansList

	followingCount, err := rt.db.UserFollowingsCount(profile.Username)
	if err != nil {
		http.Error(w, "error with following count", http.StatusBadRequest)
		return
	}
	profile.FollowingCount = followingCount

	fCount, err := rt.db.UserFollowersCount(profile.Username)
	if err != nil {
		http.Error(w, "error with follow count", http.StatusBadRequest)
		return
	}
	profile.FollowCount = fCount

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}