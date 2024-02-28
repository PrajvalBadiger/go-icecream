package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PrajvalBadiger/go-icecream/types"
)

// handle_flavour: handler to create and get flavours
func (s *Server) handle_flavour(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.create_flavour(w, r)
	}

	if r.Method == "GET" {
		return s.get_flavours(w, r)
	}

	return fmt.Errorf("Method not allowd %s", r.Method)
}

// handle_flavour_by_id: handler to create and get flavours by id
func (s *Server) handle_flavour_by_id(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.get_flavour_by_id(w, r)
	}

	if r.Method == "DELETE" {
		return s.delete_flavour(w, r)
	}

	if r.Method == "PUT" {
		return s.update_flavour(w, r)
	}

	return fmt.Errorf("Method not allowd %s", r.Method)
}

// create_flavour: handler to get list of flavours
func (s *Server) create_flavour(w http.ResponseWriter, r *http.Request) error {

	// Unmarshal json request data
	create_flavour_req := new(types.Create_flavour_request)
	if err := json.NewDecoder(r.Body).Decode(create_flavour_req); err != nil {
		return err
	}

	flavour := types.New_falvour(create_flavour_req.Name, create_flavour_req.Price)
	if err := s.store.Create_flavour(flavour); err != nil {
		return err
	}

	return write_JSON(w, http.StatusOK, flavour)
}

// get_flavours: handler to get list of flavours
func (s *Server) get_flavours(w http.ResponseWriter, r *http.Request) error {
	flavours, err := s.store.Get_flavours()
	if err != nil {
		return err
	}

	return write_JSON(w, http.StatusNotFound, flavours)
}

// get_flavour_by_id: handler to get flavour by its id
func (s *Server) get_flavour_by_id(w http.ResponseWriter, r *http.Request) error {
	// get id from store
	id, err := get_ID(r)
	if err != nil {
		return err
	}

	// get the flavour form the store
	flavour, err := s.store.Get_flavour_by_id(id)
	if err != nil {
		return write_JSON(w, http.StatusNotFound, api_error{Error: err.Error()})
	}

	// respond with the flavour
	return write_JSON(w, http.StatusOK, flavour)
}

// delete_flavour: handler to delete getflavour by its id
func (s *Server) delete_flavour(w http.ResponseWriter, r *http.Request) error {
	// get id
	id, err := get_ID(r)
	if err != nil {
		return err
	}

	// delete from the store
	if err := s.store.Delete_flavour(id); err != nil {
		return err
	}

	return write_JSON(w, http.StatusOK, map[string]int{"delete": id})
}

// update_flavour: handler to delete getflavour by its id
func (s *Server) update_flavour(w http.ResponseWriter, r *http.Request) error {
	// get id
	id, err := get_ID(r)
	if err != nil {
		return err
	}

	// Check if the flavour is present in the store
	_, err = s.store.Get_flavour_by_id(id)
	if err != nil {
		return err
	}

	// Demarshell the json
	create_flavour_req := new(types.Create_flavour_request)
	if err := json.NewDecoder(r.Body).Decode(create_flavour_req); err != nil {
		return err
	}
	flavour := types.New_falvour(create_flavour_req.Name, create_flavour_req.Price)

	// update from the store
	if err := s.store.Update_flavour(id, flavour); err != nil {
		return err
	}

	return write_JSON(w, http.StatusOK, map[string]int{"updated": id})
}
