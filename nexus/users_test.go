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

var defaultUser = User{
	UserID:    "test-user",
	FirstName: "Test",
	LastName:  "User",
	Email:     "test@user",
	Source:    "default",
	ReadOnly:  false,
	Roles:     []string{"nx-admin"},
	Status:    "active",
	Password:  "t3st-us3r",
}

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
	s := newServerWrapper(t).WithResponse(listUsersExpected).Build()
	defer s.teardown()

	users, err := s.Client().UserService.List()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestUserService_Update(t *testing.T) {
	// first, let's try updating a user which doesn't exist
	s := newServerWrapper(t).WithStatusCode(http.StatusNotFound).Build()
	assert.Error(t, s.Client().UserService.Update(defaultUser))

	// now one that does exist
	s = newServerWrapper(t).WithStatusCode(http.StatusOK).Build()
	// add the user
	assert.NoError(t, s.Client().UserService.Add(defaultUser))
	// try to update it
	assert.NoError(t, s.Client().UserService.Update(defaultUser))
}

func TestUserService_GetUserByID(t *testing.T) {
	s := newServerWrapper(t).WithResponse(adminUserResult).Build()
	defer s.teardown()

	user, err := s.Client().UserService.GetUserByID(getUserForTest())
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, getUserForTest(), user.UserID)
}

func TestUserService_AddUser(t *testing.T) {
	s := newServerWrapper(t).Build()
	defer s.teardown()

	err := s.Client().UserService.Add(User{
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
	s := newServerWrapper(t).WithResponse(validationMessage).WithStatusCode(http.StatusBadRequest).Build()
	defer s.teardown()

	err := s.Client().UserService.Add(User{
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
