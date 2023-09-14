package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"

	"os"
	"strings"
	"time"
)

var SUPPORTED_STATUS_CODES = [...]int{100, 101, 102, 103, 200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 421, 422, 423, 424, 425, 426, 428, 429, 431, 451, 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511}

func clamp(value int, minimum int, maximum int) int {
	if value < minimum {
		return minimum
	} else if value > maximum {
		return maximum
	} else {
		return value
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")

	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.StaticFile("/robots.txt", "./static/robots.txt")

	router.GET("/", handlerIndex)

	// Add the status codes
	for _, code := range SUPPORTED_STATUS_CODES {
		router.Any("/"+strconv.Itoa(code), middlewareSleep, handlerStatusCode)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func middlewareSleep(c *gin.Context) {
	sleepMs := c.Query("sleep")
	if sleepMs != "" {
		sleepMs_i, err := strconv.Atoi(sleepMs)

		if err != nil {
			panic(err)
		}

		sleepMs_i = clamp(sleepMs_i, 0, 5000)
		time.Sleep(time.Duration(sleepMs_i) * time.Millisecond)
	}

	c.Next()
}

func handlerIndex(c *gin.Context) {
	mdfile, err := os.ReadFile("./markdown/index.md")

	if err != nil {
		panic(err)
	}

	html := template.HTML(blackfriday.MarkdownCommon(mdfile))

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "statuscode.app",
		"content": html,
	})
}

func handlerStatusCode(c *gin.Context) {
	tokens := strings.Split(c.Request.URL.Path, "/")

	code := tokens[1]
	code_i, err := strconv.Atoi(code)

	if err != nil {
		panic(err)
	}

	c.String(code_i, http.StatusText((code_i)))
}
