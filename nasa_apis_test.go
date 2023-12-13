package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateExternalRequest(t *testing.T) {
	apiURL := "http://localhost:8080/sample"
	req, err := createRequest(apiURL)
	assert.NoError(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, apiURL, req.URL.String())
}

func TestManageExternalAPIRequestStatusCodeOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Send a mock status code
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"mock_key": "mock_value"}`))
	}))
	defer server.Close()

	// gin.SetMode(gin.TestMode)
	// w := httptest.NewRecorder()
	// ctx, _ := gin.CreateTestContext(w)
	// ctx.Request = &http.Request{
	// 	Header: make(http.Header),
	// 	URL:    &url.URL{Path: "http://localhost:8080/sample"},
	// }
	client := &http.Client{} //TODO: This should be mocked.
	apiURL := "http://localhost:8080/sample"
	body, status, err := manageExternalAPIRequest(client, apiURL)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to send the API request")
	assert.Empty(t, body)
	assert.Equal(t, status, 500)
}

func TestHandleExternalAPIResponseError(t *testing.T) {
	// gin.SetMode(gin.TestMode)
	// w := httptest.NewRecorder()
	// ctx, _ := gin.CreateTestContext(w)
	// ctx.Request = &http.Request{
	// 	Header: make(http.Header),
	// 	URL:    &url.URL{},
	// }
	resp := &http.Response{}
	resp.Body = http.NoBody
	resp.Status = http.StatusText(http.StatusBadGateway)
	resp.StatusCode = http.StatusBadGateway
	body, status, err := handleAPIResponse(resp)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Bad Gateway")
	assert.Empty(t, body)
	assert.Equal(t, status, 500)
}

func TestHandleExternalAPIResponseNoError(t *testing.T) {
	// gin.SetMode(gin.TestMode)
	// w := httptest.NewRecorder()
	// ctx, _ := gin.CreateTestContext(w)
	// ctx.Request = &http.Request{
	// 	Header: make(http.Header),
	// 	URL:    &url.URL{},
	// }
	resp := &http.Response{Body: io.NopCloser(bytes.NewBufferString("Hello World"))}
	resp.Status = http.StatusText(http.StatusOK)
	resp.StatusCode = http.StatusOK
	body, status, err := handleAPIResponse(resp)
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.NotEmpty(t, body)
	assert.Contains(t, body, "Hello World")
	assert.Equal(t, status, 200)
}

func TestValidateParameterValidResult(t *testing.T) {
	paramName := "earth_date"
	paramValue := "2022-09-08"
	err := validateParam(paramName, paramValue)
	assert.NoError(t, err)
}

func TestValidateParameterInvalidResult(t *testing.T) {
	paramName := "earth_date"
	paramValue := "2022-09"
	err := validateParam(paramName, paramValue)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Incorrect earth_date format.")
}

func TestValidateParameterNonExistentParameter(t *testing.T) {
	paramName := "something_else"
	paramValue := "2022-09"
	err := validateParam(paramName, paramValue)
	assert.NoError(t, err)
	assert.Nil(t, err)
}

// func TestMain(t *testing.T) {
// 	router := gin.Default()

// 	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Simulate a successful API response with mock data
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"mock_key": "mock_value"}`))
// 	}))
// 	defer mockServer.Close()

// 	apiURL := mockServer.URL + "/nasa/apod"
// 	//baseUrl = mockServer.URL

// 	req, err := http.NewRequest("GET", apiURL, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// expectedResponse := `{"data":"{\"mock_key\": \"mock_value\"}"}`
// 	// assert.Equal(t, expectedResponse, w.Body.String())

// }
