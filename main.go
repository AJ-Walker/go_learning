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
	response := gin.H{
		"status":  true,
		"data":    userList,
		"message": "User List fethced",
	}
	c.IndentedJSON(http.StatusOK, response)
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

// updateUser
func updateUserById(c *gin.Context) {
	id := c.Param("id")

	var userBody User
	var userUpdatedData User
	for i, user := range userList {
		if user.ID == id {
			if err := c.BindJSON(&userBody); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"data":    nil,
					"message": "Wrong user body",
				})
				return
			}

			userList[i].FirstName = userBody.FirstName
			userList[i].LastName = userBody.LastName
			userList[i].Email = userBody.Email

			userUpdatedData = userList[i]

		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    userUpdatedData,
		"message": "User updated successful!",
	})
}

func main() {

	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", addUser)
	router.GET("/user/:id", getUserByID)
	router.PUT("user/:id", updateUserById)

	fmt.Printf("Starting server..")
	router.Run("localhost:8080")
	fmt.Print("\n")
}
