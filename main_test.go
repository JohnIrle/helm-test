package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	// Set up the Gin router
	router := gin.Default()
	router.GET("/albums", getAlbums)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/albums", nil)
	assert.NoError(t, err)

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Invoke the getAlbums function
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var albums []album
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&albums))
	assert.Equal(t, 3, len(albums))
	assert.Equal(t, "Blue Train", albums[0].Title)
	assert.Equal(t, "Jeru", albums[1].Title)
	assert.Equal(t, "Sarah Vaughan and Clifford Brown", albums[2].Title)
}
