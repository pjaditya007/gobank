package main

import "math/rand"

type APIResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int    `json:"number"`
	Balance   int    `json:"balance"`
}

func NewAccount(firstName string, lastName string) *APIResponse {
	return &APIResponse{
		ID:        rand.Int(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int(),
		Balance:   rand.Int(),
	}
}
