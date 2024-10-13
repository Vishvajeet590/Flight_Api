package router

import (
	"Flight_Api/app/src/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	searchController := controller.NewSearchController()
	ticketController := controller.NewTicketController()

	var router = gin.Default()
	apiGroup := router.Group("/api/v1")
	{
		searchGroup := apiGroup.Group("/search")
		{
			searchGroup.POST("/location", searchController.FindFlights)
		}
		tickets := apiGroup.Group("/tickets")
		{
			tickets.POST("/book", ticketController.ReserveTickets)
			tickets.GET("/tickets", ticketController.UserTickets)
		}
	}
	return router
}
