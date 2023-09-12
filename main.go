package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"

	"os"
	"time"
)

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

	router.GET("/", func(c *gin.Context) {

		mdfile, err := os.ReadFile("./markdown/index.md")

		if err != nil {
			panic(err)
		}

		html := template.HTML(blackfriday.MarkdownCommon(mdfile))

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "Main website",
			"content": html,
		})
	})

	router.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")
		code_i, err := strconv.Atoi(code)

		if err != nil {
			panic(err)
		}

		sleepMs := c.Query("sleep")
		if sleepMs != "" {
			sleepMs_i, err := strconv.Atoi(sleepMs)

			sleepMs_i = clamp(sleepMs_i, 0, 5000)

			if err != nil {
				panic(err)
			}

			time.Sleep(time.Duration(sleepMs_i) * time.Millisecond)
		}

		c.Status(code_i)
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
