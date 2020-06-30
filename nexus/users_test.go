//     Copyright 2020 Aicura Nexus Client and/or its authors
//
//     This file is part of Aicura Nexus Client.
//
//     Aicura Nexus Client is free software: you can redistribute it and/or modify
//     it under the terms of the GNU Lesser General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     Aicura Nexus Client is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU Lesser General Public License for more details.
//
//     You should have received a copy of the GNU Lesser General Public License
//     along with Aicura Nexus Client.  If not, see <https://www.gnu.org/licenses/>.

package nexus

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var listUsersExpected = `[
	{
	  "userId": "anonymous",
	  "firstName": "Anonymous",
	  "lastName": "User",
	  "emailAddress": "anonymous@example.org",
	  "source": "default",
	  "status": "active",
	  "readOnly": false,
	  "roles": [
		"nx-anonymous"
	  ],
	  "externalRoles": []
	},
	{
	  "userId": "admin",
	  "firstName": "Administrator",
	  "lastName": "User",
	  "emailAddress": "admin@example.org",
	  "source": "default",
	  "status": "active",
	  "readOnly": false,
	  "roles": [
		"nx-admin"
	  ],
	  "externalRoles": []
	}
  ]`

var adminUserResult = `[
	{
	  "userId": "admin",
	  "firstName": "Administrator",
	  "lastName": "User",
	  "emailAddress": "admin@example.org",
	  "source": "default",
	  "status": "active",
	  "readOnly": false,
	  "roles": [
		"nx-admin"
	  ],
	  "externalRoles": []
	}
  ]`

var validationMessage = `[
	{
	  "id": "PARAMETER password",
	  "message": "may not be empty"
	}
  ]`

func TestUserService_ListUsers(t *testing.T) {
	server, teardown := newMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, _ := r.BasicAuth()
		assert.Equal(t, getUserForTest(), u)
		assert.Equal(t, getPasswordForTest(), p)
		_, err := w.Write([]byte(listUsersExpected))
		assert.NoError(t, err)
	}))
	defer teardown()

	client := newClientForTest(server)
	users, err := client.UserService.ListAllUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestUserService_GetUserByID(t *testing.T) {
	server, teardown := newMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, _ := r.BasicAuth()
		assert.Equal(t, getUserForTest(), u)
		assert.Equal(t, getPasswordForTest(), p)
		_, err := w.Write([]byte(adminUserResult))
		assert.NoError(t, err)
	}))
	defer teardown()

	client := newClientForTest(server)
	user, err := client.UserService.GetUserByID(getUserForTest())
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, getUserForTest(), user.UserID)
}

func TestUserService_AddUser(t *testing.T) {
	server, teardown := newMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, _ := r.BasicAuth()
		assert.Equal(t, getUserForTest(), u)
		assert.Equal(t, getPasswordForTest(), p)
	}))
	defer teardown()

	client := newClientForTest(server)
	err := client.UserService.AddUser(User{
		Email:     "alien@mail.com",
		UserID:    "alien",
		FirstName: "Alien",
		LastName:  "The Predator",
		Roles:     []string{"nexus-admin"},
		Source:    "default",
		Status:    "active",
		Password:  "mysupersecretpassword",
	})
	assert.NoError(t, err)
}

func TestUserService_FailtToAddUser(t *testing.T) {
	server, teardown := newMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, _ := r.BasicAuth()
		assert.Equal(t, getUserForTest(), u)
		assert.Equal(t, getPasswordForTest(), p)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(validationMessage))
		assert.NoError(t, err)
	}))
	defer teardown()

	client := newClientForTest(server)
	err := client.UserService.AddUser(User{
		Email:     "alien@mail.com",
		UserID:    "alien",
		FirstName: "Alien",
		LastName:  "The Predator",
		Roles:     []string{"nexus-admin"},
		Source:    "default",
		Status:    "active",
	})
	assert.Error(t, err)
}
