// models.user_test.go

package main

import "testing"

// Test the validity of different combinations of username/password
func TestUserValidity(t *testing.T) {
	if !isUserValid("user1", "pass1") {
		t.Fail()
	}

	if isUserValid("user2", "pass1") {
		t.Fail()
	}

	if isUserValid("user1", "") {
		t.Fail()
	}

	if isUserValid("", "pass1") {
		t.Fail()
	}

	if isUserValid("User1", "pass1") {
		t.Fail()
	}
}

// Test if a new user can be registered with valid username/password
func TestValidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()
}

// Test that a new user cannot be registered with invalid username/password
func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	// Try to register a user with a used username
	u, err := registerNewUser("user1", "pass1")

	if err == nil || u != nil {
		t.Fail()
	}

	// Try to register with a blank password
	u, err = registerNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}

// Test the function that checks for username availability
func TestUsernameAvailability(t *testing.T) {
	saveLists()

	// This username should be available
	if !isUsernameAvailable("newuser") {
		t.Fail()
	}

	// This username should not be available
	if isUsernameAvailable("user1") {
		t.Fail()
	}

	// Register a new user
	registerNewUser("newuser", "newpass")

	// This newly registered username should not be available
	if isUsernameAvailable("newuser") {
		t.Fail()
	}

	restoreLists()
}
