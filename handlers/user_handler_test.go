package handlers

import (
	"errors"
	"testing"

	servicemocks "github.com/RegiAdi/venera/mocks/services"
	"github.com/RegiAdi/venera/responses"
	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

func TestGetUserInfoHandler(t *testing.T) {

	tests := []struct {
		name              string
		responseGetDetail responses.UserResponse
		errorGetDetail    error
		wantError         bool
	}{
		{
			name: "success",
			responseGetDetail: func() responses.UserResponse {
				var userresponse responses.UserResponse
				err := faker.FakeData(&userresponse)
				if err != nil {
					t.Log(err)
				}
				return userresponse
			}(),
		},
		{
			name:              "error",
			responseGetDetail: responses.UserResponse{},
			errorGetDetail:    errors.New("User not found"),
			wantError:         true,
		},
	}

	for _, test := range tests {
		service := servicemocks.NewUserService(t)
		handler := NewUserHandler(service)

		t.Run(test.name, func(t *testing.T) {
			app := fiber.New()
			c := app.AcquireCtx(&fasthttp.RequestCtx{})

			service.On("GetUserDetail", mock.Anything).Return(test.responseGetDetail, test.errorGetDetail)

			assert.NoError(t, handler.GetUserInfoHandler(c))
		})
	}
}
