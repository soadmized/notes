package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type note struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Body string `json:"body"`
}

var notes = []note{
	{ID: 1, Title: "First note", Author: "Alex", Body: "This is the very first note to check /get and /get/id endpoints"},
}

func getNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, notes)
}

func createNote(c *gin.Context) {
	var newNote note
	if err := c.BindJSON(&newNote); err != nil {
		return
	}
	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func main() {
	router := gin.Default()
	router.POST("/get", getNotes)
	router.POST("/create", createNote)
	router.Run("localhost:8000")
}
