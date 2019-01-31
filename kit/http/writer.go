package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type VerboseResponseWriter struct {
	http.ResponseWriter
}

func NewVerboseResponseWriter(w http.ResponseWriter) *VerboseResponseWriter {
	return &VerboseResponseWriter{ResponseWriter: w}
}

// Types of errors

func (w VerboseResponseWriter) WriteInternalServerError(req *http.Request, msg string, err error) {

	if msg == "" {
		msg = "An internal server error was encountered."
	}
	w.writeError(req, msg, http.StatusInternalServerError, err.Error())
}

func (w VerboseResponseWriter) WriteBadRequest(req *http.Request, msg string, err error) {
	if msg == "" {
		msg = "Bad request."
	}
	w.writeError(req, msg, http.StatusBadRequest, err.Error())
}

func (w VerboseResponseWriter) WriteDecodeError(req *http.Request, err error) {
	w.writeError(req, "Error decoding object", http.StatusBadRequest, err.Error())
}

func (w VerboseResponseWriter) WriteJSONMarshalError(req *http.Request, obj interface{}, err error) {
	w.writeError(req, fmt.Sprintf("Error marshalling obect to JSON: %v", obj), http.StatusBadRequest, err.Error())
}

func (w VerboseResponseWriter) WriteUuidGenerationError(req *http.Request, err error) {
	w.writeError(req, "Error generating new UUID", http.StatusBadRequest, err.Error())
}

func (w VerboseResponseWriter) writeError(req *http.Request, msg string, statusCode int, err string) {

	em := fmt.Sprintf("%s error (%s) occurred due to request: %v \n. Details: %s.", statusCode, err, req, msg)

	//do some logging here
	w.WriteResponseUsingStatus(map[string]string{"error": em}, statusCode)
}

// Set default write response to 200

func (w VerboseResponseWriter) WriteResponse(obj interface{}) {

	w.WriteResponseUsingStatus(obj, http.StatusOK)
}

// Write both successful and failed requests to JSON

func (w VerboseResponseWriter) WriteResponseUsingStatus(obj interface{}, status int) {

	w.Header().Set("Content-Type", "application/json")

	json, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to marshal response to JSON."}`))
		return
	}

	w.WriteHeader(status)
	w.Write(json)
}