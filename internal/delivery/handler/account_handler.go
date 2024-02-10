package handler

import (
	"errors"
	"github.com/Federico191/freepass-2024/internal/model"
	"github.com/Federico191/freepass-2024/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountHandler struct {
	account usecase.AccountUseCaseInterface
}

func NewAccountHandler(account usecase.AccountUseCaseInterface) *AccountHandler {
	return &AccountHandler{account: account}
}

func (a AccountHandler) CreateAccount(c *gin.Context) {
	username := getUsernameFromHeader(c)

	var req model.CreateAccount

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	req.Username = username
	account, err := a.account.Create(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, account)
}

func (a AccountHandler) UpdateProfile(c *gin.Context) {
	username := getUsernameFromHeader(c)

	account, err := a.account.GetByUsername(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var req model.UpdateAccount
	req.ID = account.ID

	result, err := a.account.Update(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (a AccountHandler) DeleteAccount(c *gin.Context) {
	accountIdStr := c.Param("account_id")
	accountId, err := parseToUint(accountIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = a.account.DeleteAccount(c, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success delete account"})
}

func getUsernameFromHeader(c *gin.Context) string {
	userData, exist := c.Get("userData")
	if !exist {
		c.JSON(http.StatusInternalServerError, errorResponse(errors.New("failed to retrieve userData")))
		return ""
	}

	userDataMap, ok := userData.(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, "Failed to convert user data to map")
		return ""
	}

	username := userDataMap["username"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, errorResponse(errors.New("failed to convert")))
		return ""
	}

	return username
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
