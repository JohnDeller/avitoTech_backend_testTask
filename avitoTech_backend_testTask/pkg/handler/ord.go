package handler

import (
	"github.com/JohnDeller/avitoTech_backend_testTask"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createOrder(c *gin.Context) {
	var input avitotech.Order
	id, err := h.services.Order.CreateOrder(input.Id, input.UserId, input.Balance)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteOrder(c *gin.Context) {
	var input avitotech.Order
	Id, err := h.services.Order.DeleteOrder(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": Id,
	})
}
