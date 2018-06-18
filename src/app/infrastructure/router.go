package infrastructure

import (
	"app/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

// Router manage the routing of HTTP requests
var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSQLHandler())

	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello World!!"}) })
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
	router.GET("/users/", func(c *gin.Context) { userController.Index(c) })
	router.PUT("/users/:id", func(c *gin.Context) { userController.Update(c) })
	router.DELETE("/users/:id", func(c *gin.Context) { userController.Delete(c) })

	Router = router
}
