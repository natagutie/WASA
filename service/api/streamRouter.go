package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))
	var stream database.Stream

	stream.UserID = userID

	username, err := rt.db.GetUsername(stream.UserID)
	if err != nil {
		http.Error(w, "couldnt get the users username", http.StatusBadRequest)
		return
	}

	myStream, err := rt.db.GetStream(username, userID)
	if err != nil {
		http.Error(w, "couldnt get the users stream", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(myStream)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}