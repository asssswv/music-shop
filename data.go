package main

import (
	"errors"
)

type Storage interface {
	Create() Album
	Read() Album
	ReadOne() (Album, error)
	Update() Album
	Delete() Album
}

type MemoryStorage struct {
	albums []Album
}

func (ms MemoryStorage) Create(a Album) {
	ms.albums = append(ms.albums, a)
}

func (ms MemoryStorage) ReadOne(id string) (Album, error) {
	for _, album := range ms.albums {
		if album.ID == id {
			return album, nil
		}
	}
	return Album{}, errors.New("not found")
}

func (ms MemoryStorage) Read() []Album {
	return ms.albums
}

func (ms MemoryStorage) Update(id string, newAlbum Album) (Album, error) {
	for i := range ms.albums {
		if ms.albums[i].ID == id {

			if newAlbum.Artist == "" && newAlbum.Price == 0 && newAlbum.Title == "" {
				return Album{}, errors.New("not found")
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
	return Album{}, errors.New("not found")
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

func NewMemoryStorage() MemoryStorage {
	var albums = []Album{
		{ID: "1", Title: "Blue Train", Artist: "John", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry", Price: 17.99},
		{ID: "3", Title: "Sarah", Artist: "Sarah", Price: 39.99},
	}
	return MemoryStorage{albums: albums}
}

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
