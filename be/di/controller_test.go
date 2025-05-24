package di_test

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"legend_score/controllers/ci"
	"legend_score/di"
	"testing"
)

func TestProvideController(t *testing.T) {
	// Create a new container
	container := dig.New()

	// Provide mock dependencies that controllers need
	// This is necessary because controllers depend on usecases
	mockDependencies(container)

	// Call provideController
	di.BuildContainer(container)

	// Verify that the controllers are provided
	err := container.Invoke(func(
		authController ci.AuthController,
		userController ci.UserController,
		gameController ci.GameController,
	) {
		assert.NotNil(t, authController, "AuthController should be provided")
		assert.NotNil(t, userController, "UserController should be provided")
		assert.NotNil(t, gameController, "GameController should be provided")
	})
	assert.NoError(t, err, "Controllers should be successfully invoked")
}

// Helper function to provide mock dependencies
func mockDependencies(container *dig.Container) {
	// Mock usecases that controllers depend on
	container.Provide(func() interface{} {
		return &mockAuthUseCase{}
	})
	container.Provide(func() interface{} {
		return &mockUserUseCase{}
	})
	container.Provide(func() interface{} {
		return &mockGameUseCase{}
	})
}