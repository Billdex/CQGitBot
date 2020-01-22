package githook

import (
	"github.com/gin-gonic/gin"
	"log"
)

//handle start event
func StarHandler(c *gin.Context){
	log.Println("someone stared your repository")


	c.JSON(200, gin.H{
		"status_code": 200,
		"message": "ok",
	})
}
