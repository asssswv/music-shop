package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "bad_request"})
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album_not_found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusNoContent, album)
			albums = append(albums[:i], albums[i+1:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album_not_found"})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			_ = c.BindJSON(&albums[i])
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "not_found"})
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.POST("/albums", postAlbum)
	return router
}

func main() {
	router := getRouter()
	_ = router.Run("localhost:8080")
}
