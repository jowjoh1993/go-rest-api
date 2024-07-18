package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	//------------------------- /events --------------------------------//
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)       // only owner
	server.PUT("/events/:id", updateEvent)    // only owner
	server.DELETE("/events/:id", deleteEvent) // only owner

	//------------------------- /signup --------------------------------//
	server.POST("/signup", signup)

	//------------------------- /login ---------------------------------//
	server.POST("/login", login)

	server.GET("/users", getUsers)
}
