package repository

import (
	"errors"
	"github.com/asadbek280604/server_on_golang_example"
)

type MemoryStorage struct {
	albums []music_shop.Album
}

func NewMemoryStorage() MemoryStorage {
	var albums = []music_shop.Album{
		{ID: "1", Title: "Blue Train", Artist: "John", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry", Price: 17.99},
		{ID: "3", Title: "Sarah", Artist: "Sarah", Price: 39.99},
	}
	return MemoryStorage{albums: albums}
}

func (ms MemoryStorage) Create(a music_shop.Album) (music_shop.Album, error) {
	ms.albums = append(ms.albums, a)
	return a, nil
}

func (ms MemoryStorage) ReadOne(id string) (music_shop.Album, error) {
	for _, album := range ms.albums {
		if album.ID == id {
			return album, nil
		}
	}
	return music_shop.Album{}, errors.New("not found")
}

func (ms MemoryStorage) Read() []music_shop.Album {
	return ms.albums
}

func (ms MemoryStorage) Update(id string, newAlbum music_shop.Album) (music_shop.Album, error) {
	for i := range ms.albums {
		if ms.albums[i].ID == id {

			if newAlbum.Artist == "" && newAlbum.Price == 0 && newAlbum.Title == "" {
				return music_shop.Album{}, errors.New("not found")
			}

			if newAlbum.Artist != "" {
				ms.albums[i].Artist = newAlbum.Artist
			}

			if newAlbum.Title != "" {
				ms.albums[i].Title = newAlbum.Title
			}

			if newAlbum.Price != 0 {
				ms.albums[i].Price = newAlbum.Price
			}

			// ms.albums[i] = newAlbum
			return ms.albums[i], nil
		}
	}
	return music_shop.Album{}, errors.New("not found")
}

func (ms MemoryStorage) Delete(id string) error {
	for i, album := range ms.albums {
		if album.ID == id {
			ms.albums = append(ms.albums[:i], ms.albums[i+1:]...)
			return nil
		}
	}

	return errors.New("not found")
}
