package controller

import (
	"Flight_Api/app/src/model/dto"
	"Flight_Api/app/src/repository"
	"Flight_Api/app/src/usecase"
	"Flight_Api/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TicketController struct {
	TicketInteractor usecase.TicketInteractor
}

func NewTicketController() TicketController {
	return TicketController{
		TicketInteractor: usecase.TicketInteractor{
			SearchRepository: repository.SearchRepository{},
			TicketRepository: repository.TicketRepository{
				SearchRepository: repository.SearchRepository{},
			},
		},
	}
}

func (t TicketController) ReserveTickets(c *gin.Context) {
	var request dto.ReserveTicketRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error while Booking flights seats: %v", err)
		utils.JSONResponse(c, http.StatusBadRequest, nil, errors.New("invalid request. please try again"))
		return
	}

	res, err := t.TicketInteractor.ReserveTickets(c, request)
	if err != nil {
		log.Printf("Error while Booking flights seats: %v", err)
		utils.JSONResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, res, nil)
	return
}

func (t TicketController) UserTickets(c *gin.Context) {
	var request dto.UserTicketsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error while fetching flights tickets: %v", err)
		utils.JSONResponse(c, http.StatusBadRequest, nil, errors.New("invalid request. please try again"))
		return
	}

	res, err := t.TicketInteractor.GetUsersTickets(c, request)
	if err != nil {
		log.Printf("Error while fetching flights tickets: %v", err)
		utils.JSONResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, res, nil)
	return
}
