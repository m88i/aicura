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
)

func TestIsServerError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Non Server Error", args{
			err: &ClientError{
				HTTPStatusCode: http.StatusBadRequest,
			}},
			false,
		},
		{"Server Error", args{
			err: &ClientError{
				HTTPStatusCode: http.StatusInternalServerError,
			}},
			true,
		},
		{"Nil Error", args{
			err: nil},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsServerError(tt.args.err); got != tt.want {
				t.Errorf("IsServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAuthenticationError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Auth Error", args{
			err: &ClientError{
				HTTPStatusCode: http.StatusUnauthorized,
			}},
			true,
		},
		{"Non Auth Error", args{
			err: &ClientError{
				HTTPStatusCode: http.StatusInternalServerError,
			}},
			false,
		},
		{"Nil Error", args{
			err: nil},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAuthenticationError(tt.args.err); got != tt.want {
				t.Errorf("IsAuthenticationError() = %v, want %v", got, tt.want)
			}
		})
	}
}
