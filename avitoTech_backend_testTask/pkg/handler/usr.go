package handler

import (
	"github.com/JohnDeller/avitoTech_backend_testTask"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) updateBalance(c *gin.Context) {
	var input avitotech.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Balance < 0 {
		newErrorResponse(c, http.StatusBadRequest, "balance below zero")
		return
	}
	id, err := h.services.User.CreateUser(input)
	if err != nil {
		id, err = h.services.User.UpdateUser(input)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		goto good

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

good:
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getBalance(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	balance, err := h.services.User.GetUserBalance(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, balance)

	return
}
