package main

import "math/rand"

type Account struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Account   int64  `json:"account"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Int63(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63n(10000000),
	}
}