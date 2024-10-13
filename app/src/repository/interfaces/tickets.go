package interfaces

import (
	"Flight_Api/app/src/model/dto"
	"context"
)

type TicketRepository interface {
	GetMaxNumberOfSeats(ctx context.Context, flightID int) (int, error)
	CheckSeatAvailability(ctx context.Context, scheduleID int) (int, error)
	ReserveTickets(ctx context.Context, userID, seats, scheduleID int) (string, error)
	GetUsersTickets(ctx context.Context, userID int) ([]dto.UserTicketsResponse, error)
}
