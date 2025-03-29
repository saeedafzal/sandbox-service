package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/saeedafzal/sandbox-service/api"
)

func setup() *httptest.Server {
	mux := api.Init()
	return httptest.NewServer(mux)
}

func GET(s *httptest.Server, path string) *http.Response {
	return request(s.URL+path, http.MethodGet, nil, nil)
}

// Performs actual http request and returns response
func request(url, method string, headers http.Header, body io.Reader) *http.Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header = headers
	res, _ := http.DefaultClient.Do(req)
	return res
}
