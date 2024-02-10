package handler

import (
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	useCase usecase.UserUseCaseInterface
}

func NewUserHandler(useCase usecase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (uc UserHandler) Register(c *gin.Context) {
	var user model.UserRegister
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	createdUser, err := uc.useCase.Register(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdUser})
}

func (uc UserHandler) Login(c *gin.Context) {
	var req model.UserLogin
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	result, token, err := uc.useCase.Login(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": result, "token": token})
}
