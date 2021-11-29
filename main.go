package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type note struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Body string `json:"body"`
}

var (
	notes = []note{
	{
		ID: 1,
		Title: "First note",
		Author: "Alex",
		Body: "This is the very first note to check /get and /get/id endpoints",
	},
	}

	greet = "POST /create - makes a note \n" +
		"POST /get - get all notes\n" +
		"POST /get/:id - get note with id\n" +
		"POST /delete/:id - delete note with id"

)

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

func getNoteByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, a := range notes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note is not found"})
}

func deleteNoteByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, a := range notes {
		if a.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "note was deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note is not found"})
}

func greeting(c *gin.Context) {
	c.String(http.StatusOK, greet)
}

func main() {
	router := gin.Default()
	router.GET("/", greeting)
	router.POST("/get", getNotes)
	router.POST("/create", createNote)
	router.POST("/get/:id", getNoteByID)
	router.POST("/delete/:id", deleteNoteByID)
	err := router.Run("localhost:8000")
	if err != nil {
		panic("!!!")
	}
}
