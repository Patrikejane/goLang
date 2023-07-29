package api

import (
	"net/http"

	"com.loollab/postalapi/storage"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/postalApi", s.handleGetUserByID)
	http.HandleFunc("/postalApi/foo", s.handleGetFoo)
	http.HandleFunc("/postalApi/bar", s.handleGetBar)
	http.HandleFunc("/postalApi/name", s.handleGetName)

	return http.ListenAndServe(s.listenAddr, nil)
}
