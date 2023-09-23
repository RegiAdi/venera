package services

import (
	"errors"
	"testing"

	repositorymocks "github.com/RegiAdi/hatchet/mocks/repositories"
	"github.com/RegiAdi/hatchet/responses"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUserService(t *testing.T) {
	userRepository := repositorymocks.NewUserRepository(t)

	got := NewUserService(userRepository)

	if assert.NotNil(t, got) {
		assert.Equal(t, userRepository, got.userRepository)
	}
}

func TestGetUserDetail(t *testing.T) {
	tests := []struct {
		name         string
		apiToken     string
		response     responses.UserResponse
		errorMessage error
		wantError    bool
	}{
		{
			name:     "success",
			apiToken: faker.UUIDDigit(),
			response: func() responses.UserResponse {
				var userresponse responses.UserResponse
				err := faker.FakeData(&userresponse)
				if err != nil {
					t.Log(err)
				}
				return userresponse
			}(),
		},
		{
			name:         "error",
			apiToken:     faker.UUIDDigit(),
			response:     responses.UserResponse{},
			errorMessage: errors.New("failed get token"),
			wantError:    true,
		},
	}

	for _, test := range tests {
		userRepository := repositorymocks.NewUserRepository(t)
		service := NewUserService(userRepository)

		t.Run(test.name, func(t *testing.T) {
			userRepository.On("GetUserByAPIToken", mock.Anything).Return(test.response, test.errorMessage)

			resp, err := service.GetUserDetail(test.apiToken)
			if test.wantError {
				assert.Error(t, err, test.errorMessage)
				assert.Empty(t, resp)
			} else {
				assert.NotEmpty(t, resp, test.response)
				assert.NoError(t, err)
			}
		})
	}
}
