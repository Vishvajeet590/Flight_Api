package usecase

import (
	"Flight_Api/app/src/model/dto"
	"Flight_Api/app/src/repository/interfaces"
	"context"
	"log"
)

type SearchInteractor struct {
	SearchRepository interfaces.SearchRepository
}

func (s SearchInteractor) FindFlight(ctx context.Context, request dto.SearchRequest) ([]dto.SearchResponse, error) {
	result, err := s.SearchRepository.FindFlightByLocation(ctx, request.Source, request.Destination, request.Date)
	if err != nil {
		log.Printf("Error while searching flight by location: %v", err)
		return nil, err
	}
	return result, nil
}
