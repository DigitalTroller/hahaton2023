package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lms/models"
	"lms/utils/token"
	"net/http"
)

func ChangeImage(c *gin.Context) {
	id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the image file from the request
	imageFile, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image file"})
		return
	}

	file, err := imageFile.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image file"})
		return
	}
	defer file.Close()

	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image data"})
		return
	}

	fmt.Println(len(imageData))

	// Update the user image in the database
	err = models.UpdateUserImage(id, imageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image updated successfully"})
}
