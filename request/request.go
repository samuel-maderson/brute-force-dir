package request

import (
	"net/http"
)

type Target struct {
	Url  string
	Path string
}

func (t *Target) Action(method string, url string, path string) (*http.Response, error) {

	var resp *http.Response
	var err error
	if method == "GET" {
		resp, err = http.Get(t.Url + t.Path)
	}

	return resp, err
}
