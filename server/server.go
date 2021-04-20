package server

import (
	// "encoding/json"
	// "net/http"
	"net/http"

	"github.com/gorilla/mux"
	hnd "github.com/rishi-org-stack/loan/view"

)

func Route() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/",(hnd.Shout)).Methods("GET")
	r.HandleFunc("/",(hnd.Shout)).Methods("POST",http.MethodPatch)
	r.HandleFunc("/loans",hnd.Register).Methods(http.MethodPost)
	r.HandleFunc("/loans",hnd.GetAllLoan).Methods(http.MethodGet)
	r.HandleFunc("/loans/{id}",hnd.GetaLoan).Methods(http.MethodGet)
	r.HandleFunc("/loans/{id}",hnd.DeleteLoan).Methods(http.MethodDelete)
	r.HandleFunc("/loans/{id}",hnd.UpdateLoan).Methods(http.MethodPatch)
	return r
}
