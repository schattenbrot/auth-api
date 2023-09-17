package utils

import (
	"log"
	"net/http"
)

// SendError is the helper function for creating an error message.
// This internally then runs the writeJSON function to send the HTTP response.
func SendError(w http.ResponseWriter, logger *log.Logger, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	logger.Println(err)

	type jsonError struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	}

	theError := jsonError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}

	Send(w, logger, statusCode, theError)
}
