package albums

import (
	"errors"
	"slices"
)

type albumStore interface {
	Add(id string, album NewAlbum) (Album, error)
	Get(id string) (Album, error)
	Update(id string, album Album) error
	List() ([]Album, error)
	Remove(id string) error
}

func NewAlbumStore() albumStore {
	return &Albums{}
}

func (a *Albums) Add(id string, album NewAlbum) (Album, error) {
	newAlbum := Album{
		ID:     id,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}
	a.albums = append(a.albums, newAlbum)
	return newAlbum, nil
}

func (a *Albums) Get(id string) (Album, error) {
	for _, album := range a.albums {
		if album.ID == id {
			return album, nil
		}
	}
	return Album{}, errors.New("album not found")
}

func (a *Albums) Update(id string, album Album) error {
	for i, currentAlbum := range a.albums {
		if currentAlbum.ID == id {
			a.albums[i] =  Album{
                ID:     id,
                Title:  album.Title,
                Artist: album.Artist,
                Price:  album.Price,
            }
			return nil
		}
	}
	return errors.New("album not found")
}

func (a *Albums) List() ([]Album, error) {
    return a.albums, nil    
}

func (a *Albums) Remove(id string) error {
	for i, album := range a.albums {
        if album.ID == id {
            a.albums = slices.Delete(a.albums, i, i+1)
            return nil
        }
    }
    return errors.New("album not found")
}
