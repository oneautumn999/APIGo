package controllers

import (
	"btpn/initializers"
	"btpn/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Photostambah(c *gin.Context) {
	var Photo struct {
		Title    string
		Caption  string
		PhotoUrl string
		UserID   uint
	}

	if c.Bind(&Photo) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request tidak ditemukan",
		})

		return
	}
	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	photo := models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl, UserID: user.ID}

	result := initializers.DB.Create(&photo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal Menambah Photo",
		})

		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{})
}

func Photosdelete(c *gin.Context) {
	//id user di database
	id := c.Param("id")

	//delete user
	initializers.DB.Delete(&models.Photo{}, id)

	//respond
	c.Status(200)
}

func PhotoUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title    string
		Caption  string
		PhotoUrl string
	}

	c.Bind(&body)

	//menemukan user yang mau di update
	var photo models.Photo
	initializers.DB.First(&photo, id)

	// memilih apa yang mau di update
	initializers.DB.Model(&photo).Updates(models.Photo{
		Title:    body.Title,
		Caption:  body.Caption,
		PhotoUrl: body.PhotoUrl,
	})

	c.JSON(200, gin.H{
		"message": photo,
	})
}
