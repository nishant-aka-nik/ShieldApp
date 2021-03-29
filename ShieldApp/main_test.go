package main

import (
	"shieldapp/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Login(t *testing.T) {

	found := Login("Thor", "123")

	allowlogin := Login("Nick", "Fury")

	assert.Equal(t, false, found)

	assert.True(t, allowlogin)

}

func Login(user string, password string) bool {
	var CredentialObj []models.Credentials

	furycredential := models.Credentials{
		Name:     "Nick Fury",
		Username: "Nick", //username to login to app
		Password: "Fury", //password to login to app
		Role:     "Director",
	}

	CredentialObj = append(CredentialObj, furycredential)

	for _, eachCredentialObj := range CredentialObj {
		if eachCredentialObj.Username == user && eachCredentialObj.Password == password {
			return true
		}
	}

	return false

}
