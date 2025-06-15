package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int    `json:"balance"`
}

func NewAccount(firstName, LastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000), // Simulating ID generation
		FirstName: firstName,
		LastName:  LastName,
		Number:    int64(rand.Intn(1000000)), // Simulating a random account number
	}
}
