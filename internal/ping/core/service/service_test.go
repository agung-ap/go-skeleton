package service

import (
	"context"
	"errors"
	"go-skeleton/internal/ping/core/domain"
	"go-skeleton/internal/ping/core/port"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Test interface - only exists in test package
type MockPingRepository struct {
	mock.Mock
}

func (m *MockPingRepository) Ping(ctx context.Context, resp *domain.Ping) error {
	args := m.Called(ctx, resp)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	// Set the response if no error
	*resp = domain.Ping{Message: "ping from repository"}
	return nil
}

func TestPingService_Ping_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockPingRepository)

	// Set up mock expectations
	mockRepo.On("Ping", mock.Anything, mock.AnythingOfType("*domain.Ping")).Return(nil).Once()

	// Create service context with mock repository
	svcCtx := port.NewServiceContext(mockRepo)
	service := NewPingService(svcCtx)

	var result domain.Ping

	// Act
	err := service.Ping(ctx, &result)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "ping from repository", result.Message)
	mockRepo.AssertExpectations(t)
}

func TestPingService_Ping_RepositoryError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockPingRepository)
	expectedError := errors.New("repository error")

	// Set up mock expectations
	mockRepo.On("Ping", mock.Anything, mock.AnythingOfType("*domain.Ping")).Return(expectedError).Once()

	// Create service context with mock repository
	svcCtx := port.NewServiceContext(mockRepo)
	service := NewPingService(svcCtx)

	var result domain.Ping

	// Act
	err := service.Ping(ctx, &result)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestNewPingService(t *testing.T) {
	// Arrange
	mockRepo := new(MockPingRepository)
	svcCtx := port.NewServiceContext(mockRepo)

	// Act
	service := NewPingService(svcCtx)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, svcCtx, service.svcCtx)
}
