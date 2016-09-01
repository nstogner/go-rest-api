package models

import "errors"

var (
	// ErrNotFound is returned when no such entity exists.
	ErrNotFound = errors.New("not found")
)

// mockUserDB is a temporary map that will be user in place of a database for
// now.
var mockUserDB = map[int]User{
	0: User{
		ID:    0,
		Name:  "Bob",
		Email: "bob@example.com",
	},
	1: User{
		ID:    1,
		Name:  "Sam",
		Email: "sam@example.com",
	},
	2: User{
		ID:    2,
		Name:  "Joe",
		Email: "joe@example.com",
	},
}

// User represents a person.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// RetrieveUsers gets a slice of users. Note: currently with our mock-database
// we will never return an error... This will change later.
func RetrieveUsers(limit int) ([]User, error) {
	var users []User
	n := 0
	for _, u := range mockUserDB {
		if n >= limit {
			break
		}
		users = append(users, u)
		n++
	}
	return users, nil
}

// RetrieveUserByID gets a single user.
func RetrieveUserByID(id int) (User, error) {
	var u User
	u, found := mockUserDB[id]
	if !found {
		return u, ErrNotFound
	}
	return u, nil
}
