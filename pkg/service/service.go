package service

import (
	"github.com/asadbek280604/server_on_golang_example"
	"github.com/asadbek280604/server_on_golang_example/pkg/repository"
)

type Service interface {
	Create(a music_shop.Album) (music_shop.Album, error)
	Read() []music_shop.Album
	ReadOne(id string) (music_shop.Album, error)
	Update(id string, newAlbum music_shop.Album) (music_shop.Album, error)
	Delete(id string) error
}

func NewService() Service {
	return repository.NewPostgresStorage()
}

var Storage = NewService()
