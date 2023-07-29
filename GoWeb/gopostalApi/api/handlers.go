package api

import (
	"encoding/json"
	"net/http"

	"com.loollab/postalapi/util"
)

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	_ = util.Round2Dec(10.344434)

	json.NewEncoder(w).Encode(user)

}

func (s *Server) handleGetFoo(w http.ResponseWriter, r *http.Request) {
	user := s.store.GetFoo(5)

	_ = util.Round2Dec(10.344434)

	json.NewEncoder(w).Encode(user)

}

func (s *Server) handleGetBar(w http.ResponseWriter, r *http.Request) {
	user := s.store.GetBar(4)

	_ = util.Round2Dec(10.344434)

	json.NewEncoder(w).Encode(user)

}

func (s *Server) handleGetName(w http.ResponseWriter, r *http.Request) {
	user := s.store.GetName(4)

	_ = util.Round2Dec(10.344434)

	json.NewEncoder(w).Encode(user)

}
