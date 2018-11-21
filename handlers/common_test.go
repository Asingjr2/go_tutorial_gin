package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []tmpArticleList

// Setup portion of testing
// Need to check testing.T vs testing.M
// Testing process is interesting in gin and GO
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)		/// set mode

	os.Exit(m.Run))					// exit and run test
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *http.ResponseRecorder) bool) {
	// Creating a recorder which will hold http response info
	w := httptest.NewRecorder()

	// Creating a service that will process request and store info
	r.ServeHTTP(w, req)

	// Need to check information being returned.  What is cricial in response body.
	if !f(w) {
		t.Fail()			
	}
}

// Stores existing list in temp list for testing.
func saveLists() {
	tmpArticleList = articleList
}

// Can update/restore temp listing into global lists file.
func restoreLists() {
	artcileList = tmpArticleList
}
