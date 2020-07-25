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

type userFakeService struct{}

var usersFake = make(map[string]*User)

func (u *userFakeService) List() ([]User, error) {
	users := make([]User, 0)
	for _, user := range usersFake {
		users = append(users, *user)
	}
	return users, nil
}

func (u *userFakeService) Update(user User) error {
	usersFake[user.UserID] = &user
	return nil
}

func (u *userFakeService) GetUserByID(userID string) (*User, error) {
	return usersFake[userID], nil
}

func (u *userFakeService) Add(user User) error {
	usersFake[user.UserID] = &user
	return nil
}
