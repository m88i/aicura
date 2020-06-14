package nexus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_ListUsers(t *testing.T) {
	client := NewClient("http://localhost:8081").
		WithCredentials("admin", "admin123").
		Verbose().
		Build()
	users, err := client.UserService.ListUsers()
	assert.NoError(t, err)
	assert.True(t, len(users) > 0)
}
