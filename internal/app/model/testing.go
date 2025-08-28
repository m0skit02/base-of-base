package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "oleg.gusev12@mail.ru",
		Password: "12345678",
	}
}
