package controller

import "github.com/gin-gonic/gin"

func ShowPerson(c *gin.Context) {
	c.JSON(200, gin.H{
		"value": "ok",
	})
}
