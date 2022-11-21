package barcodecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usagifm/ikipiro/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var barcodes []models.Barcode

	models.DB.Find(&barcodes)

	c.JSON(http.StatusOK, gin.H{"barcodes": barcodes})

}

func Find(c *gin.Context) {

	var barcode models.Barcode

	barcodeNumber := c.Param("barcodeNumber")

	if err := models.DB.First(&barcode, "barcode = ?", barcodeNumber).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return

		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"barcode": barcode})

}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}
