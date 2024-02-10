package handler

import (
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CandidatePostHandler struct {
	candidatePost usecase.CandidatePostUseCaseInterface
}

func NewCandidatePostHandler(candidatePost usecase.CandidatePostUseCaseInterface) *CandidatePostHandler {
	return &CandidatePostHandler{candidatePost: candidatePost}
}

func (cp CandidatePostHandler) Create(c *gin.Context) {
	var req model.CreateCandidatePost

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	candidatePost, err := cp.candidatePost.CreatePost(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, candidatePost)
}

func (cp CandidatePostHandler) GetPost(c *gin.Context) {
	var req model.GetDelCandidatePost

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	candidatePost, err := cp.candidatePost.GetPost(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, candidatePost)
}

func (cp CandidatePostHandler) DeleteCandidatePost(c *gin.Context) {
	postIDStr := c.Param("post_id")
	candidateIDStr := c.Param("candidate_id")

	postId, err := parseToUint(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	candidateID, err := parseToUint(candidateIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	req := model.GetDelCandidatePost{
		PostId:      postId,
		CandidateId: candidateID,
	}

	err = cp.candidatePost.DeleteCandidatePost(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success delete candidate post"})
}

func parseToUint(strValue string) (uint, error) {
	uintValue, err := strconv.ParseUint(strValue, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(uintValue), nil
}
