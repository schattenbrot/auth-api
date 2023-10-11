package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// Send is the helper function for sending back an HTTP response.
func Send(w http.ResponseWriter, logger *log.Logger, status int, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		logger.Println(err)
		SendError(w, logger, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

// Send is the helper function for sending back an HTTP response.
func SendFile(w http.ResponseWriter, logger *log.Logger, file []byte) {
	fileHeader := file[:512]
	fileContentType := http.DetectContentType(fileHeader)

	w.Header().Set("Content-Type", fileContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}
