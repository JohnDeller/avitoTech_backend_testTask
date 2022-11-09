package handler

import (
	"github.com/JohnDeller/avitoTech_backend_testTask/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	usr := router.Group("/usr")
	{
		usr.POST("/user-balance", h.updateBalance)
		usr.GET("/user-balance/:id", h.getBalance)
		usr.PUT("/user-balance", h.updateBalance)
	}

	reserve := router.Group("/reserve")
	{
		reserve.POST("/reserve-funds", h.reservedFunds)
		reserve.POST("/unreserve-funds", h.unreservedFunds)
		reserve.POST("/confirm-profit", h.confirmationProfit)
	}

	orders := router.Group("/orders")
	{
		orders.POST("/create-order", h.createOrder)
		orders.POST("/delete-orders", h.deleteOrder)
	}

	return router
}
