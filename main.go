package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var userList = []User{
	{ID: "1", FirstName: "abc", LastName: "test", Email: "abc@test.com"},
	{ID: "2", FirstName: "def", LastName: "test", Email: "def@test.com"},
	{ID: "3", FirstName: "ghi", LastName: "test", Email: "ghi@test.com"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, userList)
}

func addUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new album to the slice.
	userList = append(userList, newUser)
	c.IndentedJSON(http.StatusCreated, userList)
}

func getUserByID(c *gin.Context) {
	fmt.Print(c)
	id := c.Param("id")

	for _, a := range userList {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func main() {

	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", addUser)
	router.GET("/user/:id", getUserByID)

	fmt.Printf("Starting server..")
	router.Run("localhost:8080")
	fmt.Print("\n")
}
