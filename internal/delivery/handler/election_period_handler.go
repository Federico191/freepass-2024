package handler

import (
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ElectionPeriodHandler struct {
	electionPeriod usecase.ElectionPeriodUseCaseInterface
}

func NewElectionPeriodHandler(electionPeriod usecase.ElectionPeriodUseCaseInterface) *ElectionPeriodHandler {
	return &ElectionPeriodHandler{electionPeriod: electionPeriod}
}

func (e ElectionPeriodHandler) StartElection(c *gin.Context) {
	var req model.CreateElectionPeriod
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	result, err := e.electionPeriod.Create(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, result)
}

func (e ElectionPeriodHandler) GetElection(c *gin.Context) {
	electionPeriodStr := c.Param("election_period_id")
	electPeriodId, err := parseToUint(electionPeriodStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	result, err := e.electionPeriod.Get(c, electPeriodId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, result)
}
