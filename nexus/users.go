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
	"net/url"
)

// User represents a Nexus user
type User struct {
	UserID        string   `json:"userId"`
	FirstName     string   `json:"firstName,omitempty"`
	LastName      string   `json:"lastName,omitempty"`
	Email         string   `json:"emailAddress,omitempty"`
	Source        string   `json:"source,omitempty"`
	ReadOnly      bool     `json:"readOnly,omitempty"`
	Roles         []string `json:"roles"`
	ExternalRoles []string `json:"externalRoles"`
	Status        string   `json:"status,omitempty"`
	Password      string   `json:"password,omitempty"`
}

// UserService contains all operations related to the User domain
type UserService service

// List Retrieves all users from default source
func (u *UserService) List() ([]User, error) {
	req, err := u.client.get(u.client.appendVersion("/security/users"), "")
	if err != nil {
		return nil, err
	}
	var users []User
	_, err = u.client.do(req, &users)
	return users, err
}

// Update persists a new version of an existing user on the Nexus server
func (u *UserService) Update(user User) error {
	req, err := u.client.put(u.client.appendVersion("/security/users/"+user.UserID), "", user)
	if err != nil {
		return err
	}
	_, err = u.client.do(req, nil)
	return err
}

// GetUserByID Gets the user by it's id (authentication username)
func (u *UserService) GetUserByID(userID string) (*User, error) {
	parsedURL, err := url.ParseQuery("userId=" + userID)
	if err != nil {
		return nil, err
	}
	req, err := u.client.get(u.client.appendVersion("/security/users"), parsedURL.Encode())
	if err != nil {
		return nil, err
	}
	var users []User
	_, err = u.client.do(req, &users)
	if len(users) > 0 {
		return &users[0], err
	}
	return nil, err
}

// Add adds a new user in the Nexus Server. It's worth calling `GetUserByID` to verify if the desired user is not present.
func (u *UserService) Add(user User) error {
	req, err := u.client.post(u.client.appendVersion("/security/users"), "", user)
	if err != nil {
		return err
	}
	_, err = u.client.do(req, nil)
	return err
}
