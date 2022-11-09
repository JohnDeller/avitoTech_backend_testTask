package handler

import (
	"encoding/csv"
	"github.com/JohnDeller/avitoTech_backend_testTask"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

func (h *Handler) reservedFunds(c *gin.Context) {

	var input avitotech.Reserved

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	balance, err := h.services.User.GetUserBalance(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if input.Balance < 0 {
		newErrorResponse(c, http.StatusBadRequest, "balance below zero")
		return
	}
	var _input avitotech.User
	_input.Id = input.UserId
	_input.Balance = balance - input.Balance

	if _input.Balance < 0 {
		newErrorResponse(c, http.StatusBadRequest, "balance below zero")
		return
	}

	_userId, err := h.services.User.UpdateUser(_input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	orderId, err := h.services.Order.CreateOrder(input.OrderId, input.UserId, input.Balance)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	reserveId, err := h.services.Reserve.CreateReserve(input.Id, input.UserId, input.OrderId, input.Balance)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"reserve_id": reserveId,
		"order_id":   orderId,
		"user_id":    _userId,
	})

}

func (h *Handler) unreservedFunds(c *gin.Context) {
	var input avitotech.Reserved

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Balance < 0 {
		newErrorResponse(c, http.StatusBadRequest, "balance below zero")
		return
	}

	balance, err := h.services.Reserve.GetReserveBalance(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userBalance, err := h.services.User.GetUserBalance(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var _input avitotech.User
	_input.Id = input.UserId
	_input.Balance = userBalance + balance

	_userId, err := h.services.User.UpdateUser(_input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	reserveId, err := h.services.Reserve.DeleteReserve(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         _userId,
		"reserve_id": reserveId,
		"order_id":   input.OrderId,
	})

}

func (h *Handler) confirmationProfit(c *gin.Context) {
	var input avitotech.Reserved

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Balance < 0 {
		newErrorResponse(c, http.StatusBadRequest, "balance below zero")
		return
	}

	balance, err := h.services.Reserve.GetReserveBalance(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userBalance, err := h.services.User.GetUserBalance(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var _input avitotech.User
	_input.Id = input.UserId
	_input.Balance = userBalance + balance

	_userId, err := h.services.User.UpdateUser(_input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	reserveId, err := h.services.Reserve.DeleteReserve(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	s1, s2, s3, s4 := strconv.Itoa(input.UserId), strconv.Itoa(input.OrderId), strconv.Itoa(input.Id), strconv.FormatFloat(float64(input.Balance), 'E', -1, 64)
	records := [][]string{
		{"user_id: " + s1 + " ", "order_id: " + s2 + " ", "reserve_id: " + s3 + " ", "funds: " + s4 + "\n"},
	}

	print(s1, s2, s3, s4)
	const _fileLink = "report.csv"

	file, err := os.Create(_fileLink)
	if err != nil {
		logrus.Fatalf("Cannot create CSV file:", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	//w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			logrus.Fatalf("error writing record to csv:", err)
		}
	}

	// Записываем любые буферизованные данные в подлежащий writer (стандартный вывод).

	w.Flush()
	if err := w.Error(); err != nil {
		logrus.Fatal(err)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":                   _userId,
		"reserve_id":           reserveId,
		"order_id":             input.OrderId,
		"link to the csv file": _fileLink,
	})

}
