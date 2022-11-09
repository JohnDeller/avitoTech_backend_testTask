package handler

import (
	"github.com/JohnDeller/avitoTech_backend_testTask"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createReserve(c *gin.Context) {
	var input avitotech.Reserved
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Reserve.CreateReserve(input.Id, input.UserId, input.OrderId, input.Balance)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteReserve(c *gin.Context) {
	var input avitotech.Reserved
	Id, err := h.services.Reserve.DeleteReserve(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": Id,
	})

}
