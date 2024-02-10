package handler

import (
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CandidateHandler struct {
	candidate usecase.CandidateUseCaseInterface
}

func NewCandidateHandler(candidate usecase.CandidateUseCaseInterface) *CandidateHandler {
	return &CandidateHandler{candidate: candidate}
}

func (ca CandidateHandler) Create(c *gin.Context) {
	var req model.CreateCandidate
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	result, err := ca.candidate.Create(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, result)
}

func (ca CandidateHandler) GetCandidate(c *gin.Context) {
	candidateIdStr := c.Param("candidate_id")
	candidateId, err := parseToUint(candidateIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	result, err := ca.candidate.GetById(c, &candidateId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, result)
}

func (ca CandidateHandler) DeleteCandidate(c *gin.Context) {
	candidateIdStr := c.Param("candidate_id")
	candidateId, err := parseToUint(candidateIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = ca.candidate.DeleteCandidate(c, candidateId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success delete candidate"})
}
