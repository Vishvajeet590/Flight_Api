package interfaces

import (
	"Flight_Api/app/src/model/dto"
	"context"
)

type Tickets interface {
	ReserveTickets(ctx context.Context, request dto.ReserveTicketRequest) (dto.ReserveTicketResponse, error)
	GetUsersTickets(ctx context.Context, request dto.UserTicketsRequest) ([]dto.UserTicketsResponse, error)
}
