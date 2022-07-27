package infobip

import "net/http"

// Doer is an interface for making HTTP requests.
type Doer interface {
	Do(req *http.Request) (resp *http.Response, err error)
}
