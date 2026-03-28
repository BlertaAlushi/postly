package models

import "strings"

type Register struct {
	Login
	Firstname string `json:"firstname" binding:"required,min=2,max=50"`
	Lastname  string `json:"lastname" binding:"required,min=2,max=50"`
	Email     string `json:"email" binding:"required,email"`
}

type Login struct {
	Username string `json:"username" binding:"required,alphanum,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8"`
}

type User struct {
	ID int `json:"id"`
	Register
}

type UserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (r *Register) NormalizeInputs() {
	r.Username = strings.TrimSpace(r.Username)
	r.Firstname = strings.TrimSpace(r.Firstname)
	r.Lastname = strings.TrimSpace(r.Lastname)
	r.Email = strings.TrimSpace(r.Email)
}

func (l *Login) NormalizeInputs() {
	l.Username = strings.TrimSpace(l.Username)
}
