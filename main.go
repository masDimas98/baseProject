package main

import (
	"github.com/gin-gonic/gin"
	"goTest/internal/api"
	"log"
)

// main application
func main() {

	//router initiation
	router := gin.Default()

	auth := router.Group("/auth")
	{
		//login router leads to log in controller
		auth.POST("/login", func(c *gin.Context) {
			api.Login(c)
		})
		//login router leads to log out controller
		auth.GET("/logout", func(c *gin.Context) {
			api.Logout(c)
		})
	}

	userManagement := router.Group("/userManagement")
	{
		//get router leads to get all users controller
		userManagement.GET("/getUsers", func(c *gin.Context) {
			api.GetUsers(c)
		})
		//get router leads to get user by id controller with id as param
		userManagement.GET("/getUser/:id", func(c *gin.Context) {
			api.GetUser(c)
		})
		//post router leads to create user with request body
		userManagement.POST("/createUser", func(c *gin.Context) {
			api.CreateUser(c)
		})
		//put router leads to update user with request body and id as param
		userManagement.PUT("/updateUser/:id", func(c *gin.Context) {
			api.UpdateUser(c)
		})
		//delete router leads tp delete user with id as param
		userManagement.DELETE("/deleteUser/:id", func(c *gin.Context) {
			api.DeleteUser(c)
		})
	}

	//router execution
	err := router.Run(":8080")
	//error handler
	if err != nil {
		//giving log and terminate the application
		log.Fatal("Unable to start the server")
	}
}
