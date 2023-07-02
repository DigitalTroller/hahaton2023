package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type UserImage struct {
	gorm.Model
	UserID uint   `json:"userid"`
	Image  []byte `gorm:"type:[]byte" json:"image"`
}

func UpdateUserImage(userID uint, imageData []byte) error {
	// Check if the user exists in the database
	var user User
	if err := DB.First(&user, userID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return fmt.Errorf("User not found")
		}
		log.Printf("Error checking user existence: %v", err)
		return err
	}

	// Create or update the user image in the user_image table
	userImage := UserImage{
		UserID: userID,
		Image:  imageData,
	}

	if err := DB.Where(UserImage{UserID: userID}).Assign(&userImage).FirstOrCreate(&userImage).Error; err != nil {
		log.Printf("Error updating user image: %v", err)
		return err
	}

	return nil
}
