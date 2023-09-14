package services

import (
	"testing"

	repositorymocks "github.com/RegiAdi/hatchet/mocks/repositories"
	"github.com/stretchr/testify/assert"
)

func TestNewUserService(t *testing.T) {
	userRepository := repositorymocks.NewUserRepository(t)

	got := NewUserService(userRepository)

	if assert.NotNil(t, got) {
		assert.Equal(t, userRepository, got.userRepository)
	}
}
