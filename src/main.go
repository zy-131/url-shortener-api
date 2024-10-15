package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

var urlStore = make(map[string]string)
var mu sync.Mutex

func main() {
	// Database Connection
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "url_shortener",
	}

	dsn := cfg.FormatDSN()

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}
	defer db.Close()
	fmt.Println("Connected!")

	// Gin Server Configuration and Start
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost", "127.0.0.1"})

	router.POST("/shortenURL", shortenURL)
	router.GET("/:shortURL", redirectURL)

	router.Run("localhost:8080")
}

// Shortens URL that is passed in through the request body
// Calls URL shortening method to create a unique identifier
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

// Redirects user to corresponding URL based on short URL contained in endpoint parameter
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
