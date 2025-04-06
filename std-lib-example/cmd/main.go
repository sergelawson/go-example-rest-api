package main

import (
	"net/http"

	"github.com/sergelawson/go-example-rest-api/std-lib-example/pkg/albums"
)

type homeHandler struct{}


func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page!"))
}



func main() {

	store := albums.NewAlbumStore()
	albumHandler  := albums.NewAlbumHandler(store)


	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})
	mux.Handle("/albums/", albumHandler)
	mux.Handle("/albums",  albumHandler)

	http.ListenAndServe(":8080", mux)

}
