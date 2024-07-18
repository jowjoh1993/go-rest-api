package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Multiple handlers can be specified, and they will be executed
	// sequentially from left to right
	// 	server.POST("/events", middlewares.Authenticate, createEvent)

	// Alternatively, you can group routes together
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // always run middleware for endpoints in the group
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/users", getUsers)
}
