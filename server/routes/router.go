package routes

import (
	"person-api/controller"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("v1")
	{
		person := main.Group("person")
		{
			person.GET("/", controller.ShowPerson)
		}
	}

	return router
}
