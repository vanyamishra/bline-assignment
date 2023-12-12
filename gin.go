package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// TODO: Add this to an environment variable
const apiKey = "4ecExbMPdGdE4k2WBbhK3RW3fXrqX9sA6CWtCkYm"
const apiUrlApod = "https://api.nasa.gov/planetary/apod"
const apiUrlMarsRover = "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos"

func main() {

	// Prompt the user to enter a string
	fmt.Println("Please enter the earth date in YYYY-MM-DD format to retrieve the Mars Rover Photos.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}
	earthDate := scanner.Text()
	fmt.Println("You entered: ", earthDate)

	router := gin.Default()

	router.GET("/apod", func(c *gin.Context) {
		client := &http.Client{}

		req, err := http.NewRequest("GET", fmt.Sprintf("%s?api_key=%s", apiUrlApod, apiKey), nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		//Send the request
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": string(body)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %d - %s", resp.StatusCode, resp.Status)})
		}
	})

	router.GET("/marsrover", func(c *gin.Context) {
		client := &http.Client{}

		//TODO: Accept this from the user
		//earthDate := "2015-6-3"
		req, err := http.NewRequest("GET", fmt.Sprintf("%s?earth_date=%s&api_key=%s", apiUrlMarsRover, earthDate, apiKey), nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		//Send the request
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": string(body)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %d - %s", resp.StatusCode, resp.Status)})
		}
	})
	router.Run(":8080")
}
