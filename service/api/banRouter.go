package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))
	username, err := rt.db.GetUsername(userID)
	if err != nil {
		http.Error(w, "failed getting username", http.StatusBadRequest)
		return
	}

	bUsername := ps.ByName("bUsername")

	var ban database.Bans
	ban.Username = username
	ban.BUsername = bUsername

	err = rt.db.BanUser(ban)
	if err != nil {
		http.Error(w, "failed to ban user", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")

}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))
	username, err := rt.db.GetUsername(userID)
	if err != nil {
		return
	}

	bUsername := ps.ByName("bUsername")

	var baner database.Bans
	baner.BUsername = bUsername
	baner.Username = username

	err = rt.db.UnBanUser(baner)
	if err != nil {
		http.Error(w, "failed to unban user", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func (rt *_router) checkIfBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	username := ps.ByName("username")

	bUsername := ps.ByName("bUsername")

	var baner database.Bans
	baner.Username = username
	baner.BUsername = bUsername

	isbanned, err := rt.db.IfBan(baner)
	if err != nil {
		http.Error(w, "Cannot check whether user is banned !", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(isbanned)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}