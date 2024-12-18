package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Retrieve the "id" URL parameter from the current request context, then convert it to
// an integer and return it. If the operation isn't successful, return 0 and an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

// Helper for sending JSON responses. Method takes destination http.ResponseWriter,
// HTTP status code to send, data to encode to JSON and header map for additional HTTP headers.
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// Encode data to JSON
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	js = append(js, '\n')

	// Add provided headers to response
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
