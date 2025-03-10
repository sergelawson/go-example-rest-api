package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

var albums = []album{
	{
		ID:     "933301012",
		Title:  "Troublesome world",
		Artist: "Michael B Jordan",
		Price:  44,
	},
	{
		ID:     "93440034",
		Title:  "The world of Dragon",
		Artist: "Hunter Soldier",
		Price:  50,
	},
}

func validationErrorResponse(err error) gin.H {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, strings.ToLower(e.Field())+": "+e.Tag())
	}
	return gin.H{"errors": errors}
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, validationErrorResponse(err))
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func updateAlbum(c *gin.Context) {
	albumId := c.Param("id")

	var newAlbum album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, validationErrorResponse(err))
		return
	}

	for i, currentAlbum := range albums {
		if currentAlbum.ID == albumId {
			albums[i] = newAlbum
			break
		}
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func deleteAlbum(c *gin.Context) {
	albumId := c.Param("id")

	var newAlbums []album

	for _, currentAlbum := range albums {

		if albumId == currentAlbum.ID {
			continue
		}
		newAlbums = append(newAlbums, currentAlbum)
	}

	albums = newAlbums

	c.Status(http.StatusNoContent)

}

func main() {

	router := gin.Default()

	router.GET("/albums", getAlbums)

	router.POST("/albums", postAlbum)

	router.PUT("/albums/:id", updateAlbum)

	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")

}
