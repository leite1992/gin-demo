package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	successCode = 200
)

type Demo struct {
	Test string
}

func (d *Demo) GetDemo (c *gin.Context) {
	fmt.Print(*d)
	c.JSON(successCode, gin.H{
		"message": "get demo ssssss",
	})
}

func (d *Demo) PostDemo(c *gin.Context){
	fmt.Print(*d)
	c.JSON(successCode, gin.H{
		"message": "post demo",
	})
}
