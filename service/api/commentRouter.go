package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))
	var commenter database.Comments

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		http.Error(w, "error with parsing photoId", http.StatusBadRequest)
		return
	}
	username, err := rt.db.GetUsername(userID)
	if err != nil {
		http.Error(w, "cant get username", http.StatusBadRequest)
		return
	}
	commenter.Username = username

	err = json.NewDecoder(r.Body).Decode(&commenter)
	if err != nil {
		http.Error(w, "Error in decode request", http.StatusInternalServerError)
		return
	}

	commenter.UserID = userID

	commenter.PhotoID = photoID

	doComment, err := rt.db.InsertComment(commenter)
	if err != nil {
		http.Error(w, "Error in uploading comment", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(doComment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getUserToken(r.Header.Get("Authorization"))

	commentID, err := strconv.ParseUint(ps.ByName("commentID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	var comment database.Comments

	comment.UserID = userID
	comment.CommentID = commentID

	err = rt.db.DeleteComment(comment)
	if err != nil {
		http.Error(w, "couldnt delete comment :(", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusOK)

}

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		http.Error(w, "photo ID smth wrong with it", http.StatusBadRequest)
		return
	}

	comments, err := rt.db.GetPhotoComments(photoID)
	if err != nil {
		http.Error(w, "cant get userID with username", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}