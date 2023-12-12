package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiKey = "4ecExbMPdGdE4k2WBbhK3RW3fXrqX9sA6CWtCkYm"
const baseUrl = "https://api.nasa.gov"

func main() {

	router := gin.Default()
	client := &http.Client{}

	//Adding mapping for Astronomy Picture of the Day API
	const mappingApod = "/apod"
	const apiUrlApod = baseUrl + "/planetary/apod"
	createExternalGetAPIMapping(client, router, apiUrlApod, mappingApod)

	//Adding mapping for Mars Rover Photos API
	const mappingMarsRoverPhotos = "/mars-rover-photos/:earth_date"
	const apiUrlMarsRoverPhotos = baseUrl + "/mars-photos/api/v1/rovers/curiosity/photos"
	const paramName = "earth_date"
	const promptMessage = "Please enter the earth date in YYYY-MM-DD format to retrieve the Mars Rover Photos."
	createExternalGetAPIMappingWithParameter(client, router, apiUrlMarsRoverPhotos, mappingMarsRoverPhotos, promptMessage, paramName)

	router.Run(":8080")
}

func createExternalGetAPIMapping(client *http.Client, router *gin.Engine, url string, mapping string) {
	router.GET(mapping, func(c *gin.Context) {
		//Create the GET request
		req, err := http.NewRequest("GET", fmt.Sprintf("%s?api_key=%s", url, apiKey), nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the API request"})
			return
		}

		//Send the request
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send the API request"})
			return
		}
		defer resp.Body.Close()

		//Display the response
		if resp.StatusCode == http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read the response body"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"result": string(body)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %d - %s", resp.StatusCode, resp.Status)})
		}
	})
}

func createExternalGetAPIMappingWithParameter(client *http.Client, router *gin.Engine, url string, mapping string, promptMessage string, paramName string) {
	router.GET(mapping, func(c *gin.Context) {
		//Retrieve a single parameter
		paramValue := c.Param(paramName)

		//Create the GET request
		req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s=%s&api_key=%s", url, paramName, paramValue, apiKey), nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the API request"})
			return
		}

		//Send the request
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send the API request"})
			return
		}
		defer resp.Body.Close()

		//Display the response
		if resp.StatusCode == http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read the response body"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"result": string(body)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %d - %s", resp.StatusCode, resp.Status)})
		}
	})
}
