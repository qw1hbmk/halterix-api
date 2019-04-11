package http

import (
	"fmt"
	"net/http"
	"strings"
)

// formatRequest generates ascii representation of a request
func FormatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, _ := range r.Header {
		request = append(request, fmt.Sprintf("%v", name))

	}

	// If this is a POST, add post data
	if r.Method == "POST" || r.Method == "PATCH" || r.Method == "PUT" {
		r.ParseForm()
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, " ")
}
