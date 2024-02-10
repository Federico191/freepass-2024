package test

import (
	"context"
	"errors"
	"github.com/Federico191/freepass-2024/internal/entity"
	"github.com/Federico191/freepass-2024/internal/mocks"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestUserUseCase_Register(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.UserRepositoryInterface)
		timeOut := 10 * time.Second
		useCase := usecase.NewUserUseCase(mockRepo, timeOut, nil) // MockToken tidak digunakan dalam Register
		ctx, cancel := context.WithTimeout(context.Background(), timeOut)
		defer cancel()

		req := model.UserRegister{
			Username: "riko121",
			Email:    "riko121@gmail.com",
			FullName: "riko full",
			Password: "rahasia123",
		}

		expectedUser := entity.User{
			Username: req.Username,
			Email:    req.Email,
			FullName: req.FullName,
			Password: req.Password,
		}

		mockRepo.On("GetByUsername", ctx, req.Username).Return(entity.User{}, nil)

		mockRepo.On("Create", ctx, req).Return(expectedUser, nil)

		createdAccount, err := useCase.Register(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, createdAccount)
		require.Equal(t, expectedUser, createdAccount)
	})

	t.Run("UsernameExists", func(t *testing.T) {
		mockRepo := new(mocks.UserRepositoryInterface)
		timeOut := 10 * time.Second
		useCase := usecase.NewUserUseCase(mockRepo, timeOut, nil) // MockToken tidak digunakan dalam Register
		ctx, cancel := context.WithTimeout(context.Background(), timeOut)
		defer cancel()

		req := model.UserRegister{
			Username: "riko121",
			Email:    "riko121@gmail.com",
			FullName: "riko full",
			Password: "rahasia123",
		}

		mockRepo.On("GetByUsername", ctx, req.Username).
			Return(entity.User{}, nil) // Simulate existing user

		mockRepo.On("Create", ctx, req).Return(entity.User{}, errors.New("username already exist")) // Expect error

		createdAccount, err := useCase.Register(ctx, req)
		require.Error(t, err)
		require.Empty(t, createdAccount)
		require.EqualError(t, err, "username already exist")
	})

	mock.AssertExpectationsForObjects(t)
}

func TestUserUseCase_Login(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryInterface)
	timeOut := 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	//Success case
	reqSuccess := model.UserLogin{
		Username: "riko121",
		Password: "rahasia123",
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqSuccess.Password), bcrypt.DefaultCost)

	require.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	expectSuccess := entity.User{
		Username: "riko121",
		Email:    "riko121@gmail.com",
		FullName: "riko full",
		Password: string(hashedPassword),
	}

	mockRepo.On("GetByUsername", ctx, reqSuccess.Username).Return(expectSuccess, nil)

	user, err := mockRepo.GetByUsername(ctx, reqSuccess.Username)
	require.NoError(t, err)
	assert.NotEmpty(t, user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqSuccess.Password))
	require.NoError(t, err)

	// Failed Case
	reqFail := model.UserLogin{
		Username: "salah",
		Password: "rahasia123",
	}

	expectFail := errors.New("cannot login")

	mockRepo.On("GetByUsername", ctx, reqFail.Username).Return(entity.User{}, expectFail)

	userFailed, err := mockRepo.GetByUsername(ctx, reqFail.Username)
	require.Error(t, err)
	require.Empty(t, userFailed)

	require.Equal(t, expectFail, err)

	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_Login_CreateToken(t *testing.T) {
	mockToken := new(mocks.Maker)

	expectSuccess := entity.User{
		Username: "riko121",
		Email:    "riko121@gmail.com",
		FullName: "riko full",
	}

	createdToken := "dummy token"

	mockToken.On("CreateToken", expectSuccess.Username, time.Hour*24).Return(createdToken, nil)

	token, err := mockToken.CreateToken(expectSuccess.Username, time.Hour*24)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.Equal(t, createdToken, token)

	mockToken.AssertExpectations(t)
}
