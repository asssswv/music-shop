package main

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Storage interface {
	Create(a Album) (Album, error)
	Read() []Album
	ReadOne(id string) (Album, error)
	Update(id string, newAlbum Album) (Album, error)
	Delete(id string) error
}

func NewStorage() Storage {
	return NewPostgresStorage()
}

var storage = NewStorage()
