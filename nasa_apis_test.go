package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockRequestSender is a mock implementation of RequestSender for testing purposes.
type MockRequestSender struct {
	MockResponse *http.Response
	MockErr      error
}

// SendRequest sends a mock HTTP request.
func (ms *MockRequestSender) SendRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	return ms.MockResponse, ms.MockErr
}

func TestManageExternalAPIRequestStatusOK(t *testing.T) {
	client := &http.Client{}
	apiURL := "http://localhost:8080/sample"
	mockResponse := &http.Response{
		Body:       io.NopCloser(bytes.NewBufferString("Mock API response")),
		StatusCode: http.StatusOK,
	}
	mockRequestSender := &MockRequestSender{MockResponse: mockResponse, MockErr: nil}
	body, status, err := manageAPIRequest(client, apiURL, mockRequestSender)
	assert.Nil(t, err)
	assert.NotEmpty(t, body)
	assert.Contains(t, body, "Mock API response")
	assert.Equal(t, status, 200)
}

func TestManageExternalAPIRequestStatusInternalServerError(t *testing.T) {
	client := &http.Client{} //TODO: This should be mocked.
	apiURL := "http://localhost:8080/sample"
	mockResponse := &http.Response{
		Body:       io.NopCloser(bytes.NewBufferString("Mock API response")),
		StatusCode: http.StatusOK,
	}
	mockRequestSender := &MockRequestSender{MockResponse: mockResponse, MockErr: errors.New("some error")}
	body, status, err := manageAPIRequest(client, apiURL, mockRequestSender)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "some error")
	assert.Empty(t, body)
	assert.Equal(t, status, 500)
}

func TestCreateRequest(t *testing.T) {
	apiURL := "http://localhost:8080/sample"
	req, err := createRequest(apiURL)
	assert.NoError(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, apiURL, req.URL.String())
}

func TestHandleAPIResponseError(t *testing.T) {
	resp := &http.Response{Body: io.NopCloser(bytes.NewBufferString("Hello World")), StatusCode: http.StatusBadGateway, Status: http.StatusText(http.StatusBadGateway)}
	body, status, err := handleAPIResponse(resp)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Bad Gateway")
	assert.Empty(t, body)
	assert.Equal(t, status, http.StatusBadGateway)
}

func TestHandleAPIResponseNoError(t *testing.T) {
	resp := &http.Response{Body: io.NopCloser(bytes.NewBufferString("Hello World")), StatusCode: http.StatusOK, Status: http.StatusText(http.StatusOK)}
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
