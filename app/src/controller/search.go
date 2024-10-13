package controller

import (
	"Flight_Api/app/src/model/dto"
	"Flight_Api/app/src/repository"
	"Flight_Api/app/src/usecase"
	"Flight_Api/app/src/usecase/interfaces"
	"Flight_Api/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type SearchController struct {
	SearchInteractor interfaces.Search
}

func NewSearchController() SearchController {
	return SearchController{
		SearchInteractor: usecase.SearchInteractor{
			SearchRepository: repository.SearchRepository{},
		},
	}
}

func (s SearchController) FindFlights(c *gin.Context) {
	var request dto.SearchRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error while search flights: %v", err)
		utils.JSONResponse(c, http.StatusBadRequest, nil, errors.New("invalid request. please try again"))
		return
	}
	flights, err := s.SearchInteractor.FindFlight(c, request)
	if err != nil {
		log.Printf("Error while search flights: %v", err)
		utils.JSONResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, flights, nil)
}
