package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	p4 "github.com/brettbates/p4go"
	"github.com/brettbates/p4rest/common"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

type FixesOut struct {
	Fixes []p4.Fix
}

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowFixesUnauthenticated(t *testing.T) {
	r := common.GetRouter(false)

	r.GET("/fixes", Fixes)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/fixes?change=1", nil)

	common.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {

		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		fmt.Println(string(p))
		var fixesO FixesOut
		err = json.Unmarshal(p, &fixesO)
		return err == nil && len(fixesO.Fixes) == 1 && statusOK
	})
}
