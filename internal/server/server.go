package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Server struct {
	srv http.Server
}

func copyHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		fmt.Println(err)
	}
}

func fullCopyHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("TEST", "OK")
	if _, err := w.Write(buf); err != nil {
		fmt.Println(err)
	}
}

// func(w http.ResponseWriter, r *http.Request)
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	/*
		if POST {
		} ... else if GET {
		} ... else if DELETE {
		...
		}
	*/
	// GET /users
	// POST /users {"username":"myuser"} // encoding/json
	// DELETE /users/<username> , <username>=myuser, ... // regexp
}

func New(endpoint string) *Server {
	s := &Server{}

	s.srv = http.Server{
		Addr:    endpoint,
		Handler: s,
	} // TODO: cyclic reference
	return s
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
