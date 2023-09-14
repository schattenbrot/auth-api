package utils

import (
	"encoding/json"
	"net/http"
)

// Send is the helper function for sending back an HTTP response.
func Send(w http.ResponseWriter, status int, data ...interface{}) error {
	js, err := json.Marshal(data[0])
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
