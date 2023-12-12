package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

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
	CreateExternalGetAPIMapping(client, router, apiUrlApod, mappingApod)

	//Adding mapping for Mars Rover Photos API
	const mappingMarsRoverPhotos = "/mars-rover-photos/:earth_date"
	const apiUrlMarsRoverPhotos = baseUrl + "/mars-photos/api/v1/rovers/curiosity/photos"
	const paramName = "earth_date"
	CreateExternalGetAPIMappingWithParameter(client, router, apiUrlMarsRoverPhotos, mappingMarsRoverPhotos, paramName)

	router.Run(":8080")
}

func CreateExternalGetAPIMapping(client *http.Client, router *gin.Engine, url string, mapping string) {
	router.GET(mapping, func(c *gin.Context) {
		HandleAPICall(c, client, router, url, mapping)
	})
}

func CreateExternalGetAPIMappingWithParameter(client *http.Client, router *gin.Engine, url string, mapping string, paramName string) {
	router.GET(mapping, func(c *gin.Context) {
		HandleAPICallWithParameter(c, client, router, url, mapping, paramName)
	})
}

func HandleAPICall(c *gin.Context, client *http.Client, router *gin.Engine, url string, mapping string) {
	requestURL := fmt.Sprintf("%s?api_key=%s", url, apiKey)
	req, err := CreateRequest(requestURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the API request"})
		return
	}

	resp, err := SendRequest(client, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send the API request"})
		return
	}
	defer resp.Body.Close()

	HandleAPIResponse(c, resp)
}

func HandleAPICallWithParameter(c *gin.Context, client *http.Client, router *gin.Engine, url string, mapping string, paramName string) {
	paramValue := c.Param(paramName)
	error := ValidateParam(paramName, paramValue)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	requestURL := fmt.Sprintf("%s?%s=%s&api_key=%s", url, paramName, paramValue, apiKey)
	req, err := CreateRequest(requestURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the API request"})
		return
	}

	resp, err := SendRequest(client, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send the API request"})
		return
	}
	defer resp.Body.Close()

	HandleAPIResponse(c, resp)
}

func CreateRequest(requestURL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func SendRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	return client.Do(req)
}

func HandleAPIResponse(c *gin.Context, resp *http.Response) {
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
}

func ValidateParam(paramName string, paramValue string) error {
	if paramName == "earth_date" {
		// Define a regular expression pattern for "YYYY-MM-DD"
		datePattern := "^\\d{4}-\\d{2}-\\d{2}$"
		regex := regexp.MustCompile(datePattern)
		if !regex.MatchString(paramValue) {
			return errors.New("Incorrect earth_date format. Please enter the earth_date in YYYY-MM-DD")
		}
	}
	return nil
}
