package server

import "net/http"

type HTTPServer struct {
	Name string
	*http.Server
}
