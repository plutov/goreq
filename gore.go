package gore

import (
	"net/http"
)

type Gore interface {
	Get(url string, header http.Header) (*http.Response, error)
	Post(url string, header http.Header, body []byte) (*http.Response, error)
}

type ErrorHandler func(err error)
type BeforeRequestHandler func(req *http.Request)
