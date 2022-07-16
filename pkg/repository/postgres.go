package repository

import (
	"database/sql"
	"errors"
	"github.com/asadbek280604/music-shop"
	_ "github.com/lib/pq"
	"log"
)

type PostgresStorage struct {
	db *sql.DB
}

func (p PostgresStorage) CreateSchema() error {
	_, err := p.db.Exec("create table if not exists albums(ID varchar(16) primary key NOT NULL, Title varchar(128) NOT NULL , Artist varchar(128) NOT NULL, Price decimal NOT NULL)")
	return err
}

func NewPostgresStorage() PostgresStorage {
	connStr := "user=user dbname=db password=pass sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	storage := PostgresStorage{db: db}
	err = storage.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}
	return storage
}

func (p PostgresStorage) Create(a music_shop.Album) (music_shop.Album, error) {
	_, err := p.db.Query("insert into albums(ID, Title, Artist, Price) values($1, $2, $3, $4)", a.ID, a.Title, a.Artist, a.Price)
	if err != nil {
		return music_shop.Album{}, err
	}
	return a, nil
}

func (p PostgresStorage) ReadOne(id string) (music_shop.Album, error) {
	var tempAlbum music_shop.Album
	row := p.db.QueryRow("select * from albums where id = $1", id)

	if err := row.Scan(&tempAlbum.ID, &tempAlbum.Title, &tempAlbum.Artist, &tempAlbum.Price); err != nil {
		if err == sql.ErrNoRows {
			return music_shop.Album{}, errors.New("Not found")
		}
		return music_shop.Album{}, err
	}
	return tempAlbum, nil
}

func (p PostgresStorage) Read() []music_shop.Album {
	var albums []music_shop.Album
	rows, _ := p.db.Query("select * from albums")
	defer rows.Close()
	for rows.Next() {
		var tempAlbum music_shop.Album
		rows.Scan(&tempAlbum.ID, &tempAlbum.Title, &tempAlbum.Artist, &tempAlbum.Price)
		albums = append(albums, tempAlbum)
	}
	return albums
}

func (p PostgresStorage) Update(id string, newAlbum music_shop.Album) (music_shop.Album, error) {
	result, _ := p.db.Exec("update albums set Title=$1, Artist=$2, Price=$3 where id=$4", newAlbum.Title, newAlbum.Artist, newAlbum.Price, id)
	err := handleNotFound(result)
	return newAlbum, err
}

func (p PostgresStorage) Delete(id string) error {
	result, _ := p.db.Exec("delete from albums where id=$1 ", id)
	err := handleNotFound(result)
	return err
}

func handleNotFound(result sql.Result) error {
	cnt, _ := result.RowsAffected()
	if cnt == 0 {
		return errors.New("Not Found")
	}
	return nil
}