package main

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry", Price: 17.99},
	{ID: "3", Title: "Sarah", Artist: "Sarah", Price: 39.99},
}
