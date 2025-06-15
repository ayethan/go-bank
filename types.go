package main

import "math/rand"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Number    int64
	Balance   int
}

func NewAccount(firstName, LastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000), // Simulating ID generation
		FirstName: firstName,
		LastName:  LastName,
		Number:    int64(rand.Intn(1000000)), // Simulating a random account number
	}
}
