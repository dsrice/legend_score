package di_test

import (
	"legend_score/di"
	"legend_score/repositories/ri"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func TestProvideRepository(t *testing.T) {
	// Create a new container
	container := dig.New()
	
	// Provide mock dependencies that repositories need
	// This is necessary because repositories depend on database connection
	mockDBConnection(container)
	
	// Call BuildContainer which internally calls provideRepository
	di.BuildContainer(container)
	
	// Verify that the repositories are provided
	err := container.Invoke(func(
		userRepository ri.UserRepository,
		userTokenRepository ri.UserTokenRepository,
		gameRepository ri.GameRepository,
	) {
		assert.NotNil(t, userRepository, "UserRepository should be provided")
		assert.NotNil(t, userTokenRepository, "UserTokenRepository should be provided")
		assert.NotNil(t, gameRepository, "GameRepository should be provided")
	})
	assert.NoError(t, err, "Repositories should be successfully invoked")
}

// Helper function to provide mock database connection
func mockDBConnection(container *dig.Container) {
	container.Provide(func() interface{} {
		return &mockConnection{}
	})
}

// Mock implementation of database connection
type mockConnection struct{}