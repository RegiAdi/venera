package repositories

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/RegiAdi/hatchet/models"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var apiToken = faker.UUIDDigit()

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	ctx := context.Background()

	mongodbContainer, err := mongodb.RunContainer(ctx, testcontainers.WithImage("mongo:6"))
	if err != nil {
		panic(err)
	}

	// Clean up the container
	defer func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}()

	endpoint, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		panic(err)
	}

	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(endpoint))
	if err != nil {
		panic(err)
	}

	insertToken()
	exitVal := m.Run()

	os.Exit(exitVal)
}

func insertToken() {
	db := mongoClient.Database("test")
	userCollection := db.Collection("users")
	user := models.User{
		APIToken: apiToken,
	}

	_, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetUserByAPIToken(t *testing.T) {
	tests := []struct {
		name         string
		apiToken     string
		errorMessage error
		wantError    bool
	}{
		{
			name:     "success",
			apiToken: apiToken,
		},
		{
			name:         "error",
			apiToken:     faker.UUIDDigit(),
			errorMessage: errors.New("failed get token"),
			wantError:    true,
		},
	}
	for _, test := range tests {
		service := NewUserRepository(mongoClient.Database("test"))
		assert.NotNil(t, service)

		t.Run(test.name, func(t *testing.T) {
			resp, err := service.GetUserByAPIToken(test.apiToken)
			if test.wantError {
				assert.Error(t, err, test.errorMessage)
				assert.Empty(t, resp)
			} else {
				assert.NotEmpty(t, resp, resp)
				assert.NoError(t, err)
			}
		})
	}

}
