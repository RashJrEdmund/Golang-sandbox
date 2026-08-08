package main

import "time"

type Chirp struct {
	Body string `json:"body"`
}

type ValidateChirpResponse struct {
	CleanedBody string `json:"cleaned_body"`
}

// USERS
type CreateUserRequest struct {
	Email string `json:"email"`
}

type CreateUserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CHIRPS

type CreateChirpRequest struct {
	Body   string `json:"body"`
	UserId string `json:"user_id"`
}

type CreateChirpResponse struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
