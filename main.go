package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var library = make(map[string]album)

func init() {
	for _, album := range albums {
		library[album.ID] = album
	}
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return	
	}

	library[newAlbum.ID] = newAlbum
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbums(c *gin.Context) {
	albums = maps.Values(library)
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	if album, ok := library[id]; ok {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func main() {
	router := gin.Default()
	
	albumRouter := router.Group("/albums")
	albumRouter.GET("", getAlbums)
	albumRouter.GET("/:id", getAlbumById)
	albumRouter.POST("", postAlbums)

	
	
	
	router.Run("localhost:8080")
}