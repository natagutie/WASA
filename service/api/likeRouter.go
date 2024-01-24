package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))
	var like database.Likes

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		return
	}

	like.UserID = userID
	like.PhotoID = photoID

	err = rt.db.UnlikePhoto(like) // If user already liked the photo, then we remove the like
	if err != nil {
		http.Error(w, "remove like failed", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		return
	}
	var like database.Likes
	like.UserID = userID
	like.PhotoID = photoID

	_, err = rt.db.LikePhoto(like)
	if err != nil {
		http.Error(w, "couldnt check if its liked", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}