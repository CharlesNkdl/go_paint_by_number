package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	address string
}

func NewServer() *Server {
	return &Server{
		address: ":8080",
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/health", healthHandler)

	err := http.ListenAndServe(s.address, nil)
	fmt.Println("listen and server finished")
	return err
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}
