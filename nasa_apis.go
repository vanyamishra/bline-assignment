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

	const mappingApod = "/apod"
	const apiUrlApod = baseUrl + "/planetary/apod"
	createExternalGetAPIMappingWithAPIKey(client, router, apiUrlApod, apiKey, mappingApod)

	const mappingMarsRover = "/marsrover/:earthDate"
	const apiUrlMarsRover = baseUrl + "/mars-photos/api/v1/rovers/curiosity/photos"
	const paramName = "earth_date"
	const promptMessage = "Please enter the earth date in YYYY-MM-DD format to retrieve the Mars Rover Photos."
	createExternalGetAPIMappingWithAPIKeyAndParameter(client, router, apiUrlMarsRover, apiKey, mappingMarsRover, promptMessage, paramName)

	router.Run(":8080")
}

func createExternalGetAPIMappingWithAPIKey(client *http.Client, router *gin.Engine, url string, apiKey string, mapping string) {
	router.GET(mapping, func(c *gin.Context) {
		//Create the GET request
		req, err := http.NewRequest("GET", fmt.Sprintf("%s?api_key=%s", url, apiKey), nil)
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

		//Display the response
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

func createExternalGetAPIMappingWithAPIKeyAndParameter(client *http.Client, router *gin.Engine, url string, apiKey string, mapping string, promptMessage string, paramName string) {
	router.GET(mapping, func(c *gin.Context) {
		//Accept the additional parameter via user input
		// paramValue := acceptUserInput(promptMessage)
		// if paramValue == "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "The request parameter is incorrect."})
		// 	return
		// }
		paramValue := c.Param("earthDate")
		fmt.Println(paramValue)

		//Create the GET request
		req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s=%s&api_key=%s", url, paramName, paramValue, apiKey), nil)
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
