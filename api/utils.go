package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type api_func func(http.ResponseWriter, *http.Request) error

type api_error struct {
	Error string `json:"error"`
}

// convert over handlers to http.HandlerFunc
func make_HTTP_handle_func(f api_func) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle error
			write_JSON(w, http.StatusBadRequest, api_error{Error: err.Error()})
		}
	}
}

// write_JSON: encode/Marshal any data into a json
func write_JSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

// get_ID: get a id form the http.Request
func get_ID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}
