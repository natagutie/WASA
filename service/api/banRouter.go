package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getAuth(r.Header.Get("Authorization"))

	banUsername := ps.ByName("banUsername")
	banUserID, err := rt.db.GetUserIDWithUsername(banUsername)
	if err != nil {
		return
	}
	var ban database.Ban
	ban.UserID = userID
	ban.BanUserID = banUserID

	isbanned, err := rt.db.BanCheck(ban)
	if err != nil {
		return
	}

	if !isbanned {
		err = rt.db.SetBan(ban) // If not banned then ban
		if err != nil {
			http.Error(w, "mffff helpmeeee", http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")

		message := "User banned successfully"
		w.Write([]byte(message))

	}
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getAuth(r.Header.Get("Authorization"))

	banUsername := ps.ByName("banUsername")
	banUserID, err := rt.db.GetUserIDWithUsername(banUsername)
	if err != nil {
		return
	}
	var ban database.Ban
	ban.UserID = userID
	ban.BanUserID = banUserID

	isbanned, err := rt.db.BanCheck(ban)
	if err != nil {
		return
	}
	// If already banned then we unban
	if isbanned {
		err = rt.db.SetUnBan(ban)
		if err != nil {
			http.Error(w, "mffff whyyy", http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")

		message := "User unbanned successfully"
		w.Write([]byte(message))

	}
}

func (rt *_router) isBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getAuth(r.Header.Get("Authorization"))

	banUsername := ps.ByName("banUsername")
	banUserID, err := rt.db.GetUserIDWithUsername(banUsername)
	if err != nil {
		http.Error(w, "cant find userID with username", http.StatusBadRequest)

		return
	}
	var ban database.Ban
	ban.UserID = userID
	ban.BanUserID = banUserID

	isbanned, err := rt.db.BanCheck(ban)
	if err != nil {
		http.Error(w, "Cant check if banned", http.StatusBadRequest)
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
