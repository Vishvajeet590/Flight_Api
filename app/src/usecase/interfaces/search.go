package interfaces

import (
	"Flight_Api/app/src/model/dto"
	"context"
)

type Search interface {
	FindFlight(ctx context.Context, request dto.SearchRequest) ([]dto.SearchResponse, error)
}
