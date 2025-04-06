package albums

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gosimple/slug"
)

type AlbumHandler struct {
	store albumStore
}

func NewAlbumHandler(store albumStore) *AlbumHandler {
	return &AlbumHandler{
		store: store,
	}
}

func (h *AlbumHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var (
		AlbumsRe       = regexp.MustCompile(`^/albums/*$`)
		AlbumsReWithID = regexp.MustCompile(`^/albums/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
	)

	switch {
	case r.Method == http.MethodGet && AlbumsRe.MatchString(r.URL.Path):
		h.GetAlbums(w, r)
	case r.Method == http.MethodGet && AlbumsReWithID.MatchString(r.URL.Path):
		h.GetAlbum(w, r)
	case r.Method == http.MethodPost && AlbumsRe.MatchString(r.URL.Path):
		h.CreateAlbum(w, r)
	case r.Method == http.MethodPut && AlbumsReWithID.MatchString(r.URL.Path):
		h.UpdateAlbum(w, r)
	case r.Method == http.MethodDelete && AlbumsReWithID.MatchString(r.URL.Path):
		h.DeleteAlbum(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}
}

func (a *AlbumHandler) GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := a.store.List()
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(albums)
}

func (a *AlbumHandler) GetAlbum(w http.ResponseWriter, r *http.Request) {
	albumID := r.URL.Path[len("/albums/"):]

	album, err := a.store.Get(albumID)
	if err != nil {
		NotFoundHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(album)
}

func (a *AlbumHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {

	var newAlbum NewAlbum

	if err := json.NewDecoder(r.Body).Decode(&newAlbum); err != nil {
		BadRequestHandler(w, r)
		return
	}

	AlbumID := slug.Make(newAlbum.Title)

	album, err := a.store.Add(AlbumID, newAlbum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)	
}

func (a *AlbumHandler) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	albumID := r.URL.Path[len("/albums/"):]

	var updatedAlbum Album

	if err := json.NewDecoder(r.Body).Decode(&updatedAlbum); err != nil {
		BadRequestHandler(w, r)
		return
	}

	if err := a.store.Update(albumID, updatedAlbum); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedAlbum)
}

func (a *AlbumHandler) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	albumID := r.URL.Path[len("/albums/"):]

	if err := a.store.Remove(albumID); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("204 No Content"))
}


func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("404 Not Found"))
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 Bad Request"))
}