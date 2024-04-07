package request

import (
	"net/http"
	"strconv"
)

type RequestOption func(*http.Request) error

func WithPagination(skip int, take int) RequestOption {
	return func(req *http.Request) error {
		req.URL.Query().Set("skip", strconv.Itoa(skip))
		req.URL.Query().Set("take", strconv.Itoa(take))
		return nil
	}
}
