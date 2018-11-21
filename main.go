// Contains router information and overall controller logic
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")		// Creates storage container for all included elements
	
	initializeRoutes()

	router.Run()
	
}

// Will need to build app in cmd ln (go build -o app) and then run as ./app
/* 
router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",			
			gin.H{"title": "Home Page",},		// Association of template variable!!!
		)
	})
*/
