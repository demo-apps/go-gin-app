// middleware.auth_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Test the ensureLoggedIn middleware when the user is not logged in
func TestEnsureLoggedInUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(false), ensureLoggedIn(), func(c *gin.Context) {
		// Use the setLoggedIn middleware to set the is_logged_in flag to false
		// Since we aren't logged in, this handler should not be executed.
		// If it is, then the ensureLoggedIn middleware isn't working as expected
		t.Fail()
	})

	// Use the helper method to execute process the request and test
	// the HTTP status code
	testMiddlewareRequest(t, r, http.StatusUnauthorized)
}

// Test the ensureLoggedIn middleware when the user is logged in
func TestEnsureLoggedInAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(true), ensureLoggedIn(), func(c *gin.Context) {
		// Use the setLoggedIn middleware to set the is_logged_in flag to true
		// Since we are logged in, this handler should be executed.
		c.Status(http.StatusOK)
	})

	// Use the helper method to execute process the request and test
	// the HTTP status code
	testMiddlewareRequest(t, r, http.StatusOK)
}

// Test the ensureNotLoggedIn middleware when the user is logged in
func TestEnsureNotLoggedInAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(true), ensureNotLoggedIn(), func(c *gin.Context) {
		// Use the setLoggedIn middleware to set the is_logged_in flag to true
		// Since we are logged in, this handler should not be executed.
		// If it is, then the ensureNotLoggedIn middleware isn't working as expected
		t.Fail()
	})

	// Use the helper method to execute process the request and test
	// the HTTP status code
	testMiddlewareRequest(t, r, http.StatusUnauthorized)
}

// Test the ensureNotLoggedIn middleware when the user is not logged in
func TestEnsureNotLoggedInUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(false), ensureNotLoggedIn(), func(c *gin.Context) {
		// Use the setLoggedIn middleware to set the is_logged_in flag to false
		// Since we are not logged in, this handler should be executed.
		c.Status(http.StatusOK)
	})

	// Use the helper method to execute process the request and test
	// the HTTP status code
	testMiddlewareRequest(t, r, http.StatusOK)
}

// Test the setUserStatus middleware when the user is logged in
func TestSetUserStatusAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setUserStatus(), func(c *gin.Context) {
		// as the token cookie was set, the "is_logged_in" should have been set
		// to true by the setUserStatus middleware
		loggedInInterface, exists := c.Get("is_logged_in")
		if !exists || !loggedInInterface.(bool) {
			t.Fail()
		}
	})

	// Create a response recorder
	w := httptest.NewRecorder()

	// Set the cookie
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)
}

// Test the setUserStatus middleware when the user is not logged in
func TestSetUserStatusUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setUserStatus(), func(c *gin.Context) {
		// as the token cookie was not set, the "is_logged_in" should have been set
		// to false by the setUserStatus middleware
		loggedInInterface, exists := c.Get("is_logged_in")
		if exists && loggedInInterface.(bool) {
			t.Fail()
		}
	})

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create a request to send to the above route (without any cookies)
	req, _ := http.NewRequest("GET", "/", nil)

	// Create the service and process the above request.
	r.ServeHTTP(w, req)
}

// This is a middleware that will set the value of "is_logged_in" to
// true or false depending on the value passed in. This is used only for testing
func setLoggedIn(b bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}
