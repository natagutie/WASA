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
	userID := getAuth(r.Header.Get("Authorization"))

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	var comment database.Comment
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Error decoding", http.StatusInternalServerError)
		return
	}

	comment.UserID = userID

	comment.PhotoID = photoID

	username, err := rt.db.GetUsernameWithUserID(comment.UserID)
	if err != nil {
		http.Error(w, "cant get userID with username", http.StatusBadRequest)
		return
	}
	comment.Username = username

	commenter, err := rt.db.SetComment(comment)
	if err != nil {
		http.Error(w, "Error saving comment", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(commenter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := getAuth(r.Header.Get("Authorization"))

	commentID, err := strconv.ParseUint(ps.ByName("commentID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	var comment database.Comment

	comment.UserID = userID
	comment.CommentID = commentID

	err = rt.db.RemoveComment(comment)
	if err != nil {
		http.Error(w, "Error deleting commentt", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusOK)

}

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	comments, err := rt.db.GetComments(photoID)
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
