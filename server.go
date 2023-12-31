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
	store      storage
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func NewAPIServer(listenAddr string, store storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
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

	router.HandleFunc("/accounts", MakeAPIHandler(a.handleAccount))
	router.HandleFunc("/accounts/{id}", MakeAPIHandler(a.handleGETAccount))

	log.Println("JSON API Server is running on port : ", a.listenAddr)

	http.ListenAndServe(a.listenAddr, router)
}

func (a *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return a.handleGetAccounts(w, r)
	case "POST":
		return a.handleCreateAccount(w, r)
	case "DELETE":
		return a.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("this method does not support : %s", r.Method)
	}
}

func (a *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := a.store.GetAccounts()
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, accounts)
}

func (a *APIServer) handleGETAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	if id == "" {
		return fmt.Errorf("you must write account id")
	}

	account := NewAccount("jj", "test")
	return WriteJson(w, http.StatusOK, account)
}

func UnMarshall[T any](r *http.Request) (*T, error) {
	requestInfo := new(T)

	if err := json.NewDecoder(r.Body).Decode(requestInfo); err != nil {
		return nil, err
	}

	return requestInfo, nil
}

func (a *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	requestInfo, err := UnMarshall[CreateAccountRequest](r)

	if err != nil {
		return err
	}

	account := NewAccount(requestInfo.FirstName, requestInfo.LastName)

	if err := a.store.CreateAccount(account); err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, account)
	return nil
}

func (a *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (a *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {

	return nil
}
