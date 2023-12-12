package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Add this to an environment variable
const apiKey = "4ecExbMPdGdE4k2WBbhK3RW3fXrqX9sA6CWtCkYm"
const apiUrlApod = "https://api.nasa.gov/planetary/apod"
const apiUrlMarsRover = "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos"

func main() {

	router := gin.Default()

	callExternalAPIWithParameters(router)

	callExternalAPI(router)

	router.Run(":8080")
}

func callExternalAPI(router *gin.Engine) {
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
}

func callExternalAPIWithParameters(router *gin.Engine) {
	router.GET("/marsrover", func(c *gin.Context) {
		client := &http.Client{}

		earthDate := AcceptUserInput("Please enter the earth date in YYYY-MM-DD format to retrieve the Mars Rover Photos.")

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
}