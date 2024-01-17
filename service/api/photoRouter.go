package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/natagutie/WASA/service/api/reqcontext"
	"github.com/natagutie/WASA/service/database"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var pic database.Photo

	userID := getAuth(r.Header.Get("Authorization"))

	photo, err := io.ReadAll(r.Body) //reads image file from requestbody
	if err != nil {
		http.Error(w, "Invalid photo file", http.StatusBadRequest)
		return
	}

	pic.UserID = userID
	username, err := rt.db.GetUsernameWithUserID(pic.UserID)
	if err != nil {
		http.Error(w, "cant get username with userID", http.StatusBadRequest)
		return
	}
	pic.Username = username
	nowtime := time.Now()
	pic.Date = nowtime.Format("2006-01-02 15:04:05")
	pic.Photo = photo
	pic.CommentNum = 0
	pic.LikesNum = 0

	err = rt.db.SetPic(pic)
	if err != nil {
		http.Error(w, "Error inserting picture", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(pic)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var pic database.Photo

	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	userID := getAuth(r.Header.Get("Authorization"))

	pic.UserID = userID
	pic.PhotoID = photoID

	err = rt.db.RemovePic(pic)
	if err != nil {
		http.Error(w, "Error saving comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userID := getAuth(r.Header.Get("Authorization"))
	photoUserID := ps.ByName("photoUserID")
	var photo database.Photo
	photo.UserID = userID
	photo.PhotoUserID = photoUserID

	pics, err := rt.db.GetPhotos(photo)
	if err != nil {
		http.Error(w, "cant get photo list", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(pics)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
