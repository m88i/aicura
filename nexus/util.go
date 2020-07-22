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

// NewString copies the argument string and returns a pointer to it
func NewString(str string) *string {
	return &str
}

// NewBool is an easy access to the address of bool types
func NewBool(bln bool) *bool {
	return &bln
}

// NewRepositoryType easy access to a pointer for RepositoryType
func NewRepositoryType(rType RepositoryType) *RepositoryType {
	return &rType
}

// NewRepositoryFormat easy access to a pointer for RepositoryFormat
func NewRepositoryFormat(format RepositoryFormat) *RepositoryFormat {
	return &format
}
