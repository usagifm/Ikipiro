package main

import (
	"github.com/gin-gonic/gin"
	"github.com/usagifm/ikipiro/controllers/barcodecontroller"
	"github.com/usagifm/ikipiro/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/barcode/:barcodeNumber", barcodecontroller.Find)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from the server side !",
		})
	})
	r.Run()
}
