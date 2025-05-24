package di_test

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"legend_score/repositories/ri"
	"legend_score/usecases"
	"legend_score/usecases/ui"
	"testing"
)

func TestProvideUseCase(t *testing.T) {
	// Create a new container
	container := dig.New()

	// Provide mock dependencies that usecases need
	// This is necessary because usecases depend on repositories
	mockRepositories(container)

	// Directly provide the usecase constructors
	container.Provide(usecases.NewAuthUseCase)
	container.Provide(usecases.NewUserUseCase)
	container.Provide(usecases.NewGameUseCase)

	// Verify that the usecases are provided
	err := container.Invoke(func(
		authUseCase ui.AuthUseCase,
		userUseCase ui.UserUseCase,
		gameUseCase ui.GameUseCase,
	) {
		assert.NotNil(t, authUseCase, "AuthUseCase should be provided")
		assert.NotNil(t, userUseCase, "UserUseCase should be provided")
		assert.NotNil(t, gameUseCase, "GameUseCase should be provided")
	})
	assert.NoError(t, err, "UseCases should be successfully invoked")
}

// Helper function to provide mock repositories
func mockRepositories(container *dig.Container) {
	// Provide mock repositories
	container.Provide(func() ri.UserRepository {
		return &mockUserRepository{}
	})
	container.Provide(func() ri.UserTokenRepository {
		return &mockUserTokenRepository{}
	})
	container.Provide(func() ri.GameRepository {
		return &mockGameRepository{}
	})
}