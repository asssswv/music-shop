package handler

import (
	"net/http"

	"github.com/asadbek280604/music-shop"
	"github.com/asadbek280604/music-shop/pkg/service"
	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.Storage.Read())
}

func postAlbum(c *gin.Context) {
	var newAlbum music_shop.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "bad_request"})
		return
	}

	if newAlbum.ID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "bad_request"})
	}

	service.Storage.Create(newAlbum)
	//albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := service.Storage.ReadOne(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album_not_found"})
	}

	c.IndentedJSON(http.StatusOK, album)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	err := service.Storage.Delete(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album_not_found"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, music_shop.Album{})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var newAlbum music_shop.Album
	_ = c.BindJSON(&newAlbum)

	album, err := service.Storage.Update(id, newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "not_found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.POST("/albums", postAlbum)
	return router
}