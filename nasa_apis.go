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
	router.GET("nasa/apod", func(c *gin.Context) {
		requestURL := fmt.Sprintf("%s?api_key=%s", baseUrl+"/planetary/apod", apiKey)
		body, status, err := manageExternalAPIRequest(client, requestURL)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": body})
		}
	})

	//Adding mapping for Mars Rover Photos API
	router.GET("nasa/mars-rover-photos/:earth_date", func(c *gin.Context) {
		paramName := "earth_date"
		paramValue := c.Param(paramName)
		error := validateParam(paramName, paramValue)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			return
		}
		requestURL := fmt.Sprintf("%s?%s=%s&api_key=%s", baseUrl+"/mars-photos/api/v1/rovers/curiosity/photos", paramName, paramValue, apiKey)
		body, status, err := manageExternalAPIRequest(client, requestURL)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": body})
		}
	})

	router.Run(":8080")
}

func manageExternalAPIRequest(client *http.Client, requestURL string) (string, int, error) {
	req, err := createRequest(requestURL)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	//Send request
	resp, err := sendRequest(client, req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, status, error := handleAPIResponse(resp)
	return body, status, error
}

func sendRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Failed to send the API request")
	}
	return resp, nil
}

func createRequest(requestURL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, errors.New("Failed to create the API request")
	}
	return req, nil
}

func handleAPIResponse(resp *http.Response) (string, int, error) {
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", http.StatusInternalServerError, errors.New("Failed to read the response body")
		}
		return string(body), http.StatusOK, nil
	} else {
		return "", http.StatusInternalServerError, errors.New(fmt.Sprintf("Error: %d - %s", resp.StatusCode, resp.Status))
	}
}

func validateParam(paramName string, paramValue string) error {
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
