package interfaces

import (
	"Flight_Api/app/src/model/dto"
	"context"
)

type SearchRepository interface {
	FindFlightByLocation(ctx context.Context, source, destination, date string) ([]dto.SearchResponse, error)
	FindScheduleByID(context context.Context, scheduleID int) (dto.SearchResponse, error)
}
