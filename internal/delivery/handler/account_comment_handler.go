package handler

import (
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountCommentHandler struct {
	accountComment usecase.AccountCommentUseCaseInterface
	account        usecase.AccountUseCaseInterface
}

func NewAccountCommentHandler(accountComment usecase.AccountCommentUseCaseInterface,
	account usecase.AccountUseCaseInterface) *AccountCommentHandler {
	return &AccountCommentHandler{
		accountComment: accountComment,
		account:        account,
	}
}

func (ac AccountCommentHandler) CreateComment(c *gin.Context) {
	var req model.CreateAccountComment
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accountComment, err := ac.accountComment.Create(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, accountComment)
}

func (ac AccountCommentHandler) UpdateComment(c *gin.Context) {
	username := getUsernameFromHeader(c)

	account, err := ac.account.GetByUsername(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	var req model.UpdateAccountComment
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	postIDStr := c.Param("post_id")
	candidateIDStr := c.Param("candidate_id")

	postId, err := parseToUint(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	candidateId, err := parseToUint(candidateIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	req.CandidateId = candidateId
	req.AccountId = account.ID
	req.PostId = postId

	result, err := ac.accountComment.Update(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, result)
}

func (ac AccountCommentHandler) GetComment(c *gin.Context) {
	username := getUsernameFromHeader(c)

	account, err := ac.account.GetByUsername(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	var req model.GetDelAccountComment
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	req.AccountId = account.ID

	result, err := ac.accountComment.Get(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, result)
}

func (ac AccountCommentHandler) DeleteComment(c *gin.Context) {
	username := getUsernameFromHeader(c)

	account, err := ac.account.GetByUsername(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	var req model.GetDelAccountComment
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	req.AccountId = account.ID

	err = ac.accountComment.Delete(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success delete comment"})
}
