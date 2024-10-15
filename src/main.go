package main

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var urlStore = make(map[string]string)
var mu sync.Mutex

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost", "127.0.0.1"})

	router.POST("/shortenURL", shortenURL)
	router.GET("/:shortURL", redirectURL)

	router.Run("localhost:8080")
}

func shortenURL(c *gin.Context) {
	var url struct {
		LongURL string `json:"long_url" binding:"required"`
	}

	if err := c.BindJSON(&url); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No URL passed in"})
	}

	if !strings.HasPrefix(url.LongURL, "http://") && !strings.HasPrefix(url.LongURL, "https://") {
		url.LongURL = "https://" + url.LongURL
	}

	mu.Lock()
	shortURL := GenerateShortURL()
	urlStore[shortURL] = url.LongURL
	mu.Unlock()

	c.IndentedJSON(http.StatusCreated, shortURL)
}

func redirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")

	mu.Lock()
	longURL, exists := urlStore[shortURL]
	mu.Unlock()

	if !exists {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Short URL not found"})
		return
	}

	c.Redirect(http.StatusFound, longURL)
}
