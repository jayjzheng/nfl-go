package api

import "net/http"

type Client struct {
	*http.Client
}
