package request

import "net/http"

type Interface interface {
	// Endpoint returns the underlying base URL or URL format.
	Endpoint() string

	// Request performs the actual HTTP request.
	//
	//     inp[0] the new HTTP request to be sent
	//     inp[1] the response structure to receive, if any
	//
	Request(*http.Request, any) error
}
