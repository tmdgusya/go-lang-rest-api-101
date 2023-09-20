package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
}

func MakeAPIHandler(f apiFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := f(writer, request); err != nil {
			// handle error
			WriteJson(writer, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (a *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", MakeAPIHandler(a.handleAccount))

	log.Println("JSON API Server is running on port : ", a.listenAddr)

	http.ListenAndServe(a.listenAddr, router)
}

func (a *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return a.handleGETAccount(w, r)
	case "POST":
		return a.handleCreateAccount(w, r)
	case "DELETE":
		return a.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("this method does not support : %s", r.Method)
	}
}

func (a *APIServer) handleGETAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("jj", "test")
	return WriteJson(w, http.StatusOK, account)
}

func (a *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (a *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (a *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {

	return nil
}
