package handler

import (
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountVoteHandler struct {
	accountVote usecase.AccountVoteInterface
}

func NewAccountVoteHandler(accountVote usecase.AccountVoteInterface) *AccountVoteHandler {
	return &AccountVoteHandler{accountVote: accountVote}
}

func (a AccountVoteHandler) Vote(c *gin.Context) {
	var req model.CreateAccountVote
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	result, err := a.accountVote.Vote(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, result)
}
