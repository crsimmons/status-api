package handlers

import (
	"net/http"
)

// GetHealth checks the health of the app
func GetHealth(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(200)
}
