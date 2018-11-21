package main 

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Creating simulate route request.
	req, ) := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *http.ResponseRecorder) bool {
		// Testing response status
		statusOK := w.Code == http.statusOK

		// Test page title
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		// Return results of status check and title check.
		// Can review third party libraries to use with gin that parse HTML.
		return statusOK && pageOK
	})
}
