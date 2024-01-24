package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)


func (rt *_router) logUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")
	var user database.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Cant decode username login", http.StatusBadRequest)
		return
	}
	checkUser, err := rt.db.UserExist(user.Username)
	if err != nil {
		http.Error(w, "Cant check if the user exists", http.StatusBadRequest)
		return
	}

	if checkUser {
		userID, err := rt.db.GetUserID(user.Username)
		if err != nil {
			http.Error(w, "Cant get userID", http.StatusBadRequest)
			return
		}
		user.UserID = userID
	} else {
		userID, err := rt.db.LogUser(user)
		if err != nil {
			http.Error(w, "Couldnt create account", http.StatusConflict)
			return
		}
		user.UserID = userID

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(userID)
		if err != nil {
			http.Error(w, "Couldnt send UserID response", http.StatusConflict)
			return
		}
	}

}

func (rt *_router) checkUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	username := ps.ByName("username")

	ifUser, err := rt.db.UserExist(username)

	if err != nil {
		http.Error(w, "Cannot check if the user exists.", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ifUser)
	if err != nil {
		http.Error(w, "couldnt encode ifUser response", http.StatusConflict)
		return
	}

}

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))
	username := ps.ByName("username")

	_, err := rt.db.SetNewUsername(username, userID)
	if err != nil {
		http.Error(w, "Couldnt change username", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}