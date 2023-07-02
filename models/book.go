package models

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"lms/utils/token"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Book represents data about a book.
type Book struct {
	gorm.Model
	Id          string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func GetBookByID(c *gin.Context) {

	bookID := c.Param("id") // Replace with the ID of the book you want to query

	// Construct the API URL
	apiURL := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes/%s", bookID)

	// Make the API request
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the JSON response
	var bookData interface{}
	err = json.Unmarshal(body, &bookData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookData})

}

func GetAllBooks(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	joinedString := strings.Join(u.Interests, " ")
	params := url.Values{}
	params.Add("q", joinedString)
	params.Add("key", "AIzaSyD0NiJzjDkYBVZ78FSeTqsCFDMS534kZX4") // Replace with your Google Books API key

	// Retrieve the skip and limit parameters from query parameters
	skipStr := c.DefaultQuery("skip", "0")
	limitStr := c.DefaultQuery("limit", "10")

	// Convert the skip and limit parameters to integers
	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skip value"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	params.Add("startIndex", strconv.Itoa(skip))
	params.Add("maxResults", strconv.Itoa(limit))
	params.Set("langRestrict", "en")
	params.Set("orderBy", "relevance")

	fmt.Println(limit)

	// Construct the URL with parameters
	apiURL := "https://www.googleapis.com/books/v1/volumes?" + params.Encode()

	fmt.Println(apiURL)

	// Make the API request
	response, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Parse the JSON response
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the JSON data
	c.JSON(http.StatusOK, gin.H{"data": data})
}