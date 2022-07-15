package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, storage.Read())
}

func postAlbum(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "bad_request"})
		return
	}

	if newAlbum.ID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "bad_request"})
	}

	storage.Create(newAlbum)
	//albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := storage.ReadOne(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album_not_found"})
	}

	c.IndentedJSON(http.StatusOK, album)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	err := storage.Delete(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album_not_found"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, Album{})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var newAlbum Album
	_ = c.BindJSON(&newAlbum)

	album, err := storage.Update(id, newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "not_found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func getRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.POST("/albums", postAlbum)
	return router
}
