package albums

import (
	"net/http"
	"regexp"
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
		AlbumsRe       = regexp.MustCompile(`^/recipes/*$`)
		AlbumsReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
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

func (a *AlbumHandler) GetAlbums(w http.ResponseWriter, r *http.Request) {}

func (a *AlbumHandler) GetAlbum(w http.ResponseWriter, r *http.Request) {}

func (a *AlbumHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {}

func (a *AlbumHandler) UpdateAlbum(w http.ResponseWriter, r *http.Request) {}

func (a *AlbumHandler) DeleteAlbum(w http.ResponseWriter, r *http.Request) {}
