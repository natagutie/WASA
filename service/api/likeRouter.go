package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getAuth(r.Header.Get("Authorization"))

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64) // Converting string to integer (uint64)
	if err != nil {
		return
	}
	var like database.Like
	like.UserID = userID
	like.PhotoID = photoID

	liker, err := rt.db.SetLike(like) // If its not liked then we set the like in database
	if err != nil {
		http.Error(w, "couldnt check if its liked", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(liker.LikeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getAuth(r.Header.Get("Authorization"))
	var like database.Like

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		return
	}

	like.UserID = userID
	like.PhotoID = photoID

	isLiked, err := rt.db.IsLiked(like)
	if err != nil {
		http.Error(w, "couldnt check if its liked", http.StatusBadRequest)
		return
	}

	if isLiked {
		err = rt.db.RemoveLike(like) // If user already liked the photo, then we remove the like
		if err != nil {
			http.Error(w, "remove like failed", http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)

	}

}
