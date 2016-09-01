package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

// User will be defined here for now.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// API returns back a handler with all REST API routes registered.
func API() http.Handler {
	rtr := httprouter.New()
	rtr.GET("/users/:id", getSingleUser)
	return rtr
}

func getSingleUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Let's log the request
	logrus.WithField("path", r.URL.Path).Info("new request")

	// Validate that the id in the request path is an integer
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		respondError(w, "id must be an integer", http.StatusBadRequest)
		return
	}

	// For now, we will just send back a hardcoded user
	u := User{
		ID:   id,
		Name: "bob",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// respondError sends an error message as JSON.
func respondError(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"error": err,
	})
}
