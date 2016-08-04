// models.user.go

package main

import "errors"

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real application, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
var userList = []user{
	user{Username: "user1", Password: "pass1"},
	user{Username: "user2", Password: "pass2"},
	user{Username: "user3", Password: "pass3"},
}

// Register a new user with the given username and password
func registerNewUser(username, password string) (*user, error) {
	return nil, errors.New("placeholder error")
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
	return false
}
