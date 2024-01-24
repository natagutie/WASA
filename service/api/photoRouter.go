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

	userID := getUserToken(r.Header.Get("Authorization"))

	photo, err := io.ReadAll(r.Body)
	var picture database.Photos

	picture.Photo = photo
	if err != nil {
		http.Error(w, "Something wrong with reading photo file", http.StatusBadRequest)
		return
	}

	picture.UserID = userID
	username, err := rt.db.GetUsername(picture.UserID)
	if err != nil {
		http.Error(w, "cant get username with userID", http.StatusBadRequest)
		return
	}
	picture.Username = username
	nowtime := time.Now()
	picture.Date = nowtime.Format("2006-01-02 15:04:05")
	picture.Photo = photo

	picture.PUsername = username

	err = rt.db.UploadPhoto(picture)
	if err != nil {
		http.Error(w, "Error uploading photo into database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(picture)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var picture database.Photos
	photoID, err := strconv.ParseUint(ps.ByName("photoID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}
	userID := getUserToken(r.Header.Get("Authorization"))

	picture.UserID = userID
	picture.PhotoID = photoID

	err = rt.db.DeletePhoto(picture)
	if err != nil {
		http.Error(w, "Couldnt delete photo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userID := getUserToken(r.Header.Get("Authorization"))
	pUsername := ps.ByName("photoUsername")
	var p database.Photos
	p.PUsername = pUsername

	pictures, err := rt.db.GetPhotos(p, userID)
	if err != nil {
		http.Error(w, "cannot get the photo list", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(pictures)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}