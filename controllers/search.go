package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func SearchQueryBook(c *gin.Context) {
	search_query := c.Query("q")
	fmt.Println(search_query)

	params := url.Values{}
	params.Add("q", search_query)
	params.Add("key", "AIzaSyD0NiJzjDkYBVZ78FSeTqsCFDMS534kZX4")
	apiURL := "https://www.googleapis.com/books/v1/volumes?"

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
	params.Add("langRestrict", "en")
	params.Add("orderBy", "relevance")
	apiURL += params.Encode()
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
