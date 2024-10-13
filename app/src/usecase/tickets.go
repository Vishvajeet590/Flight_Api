package usecase

import (
	"Flight_Api/app/src/model/dto"
	"Flight_Api/app/src/repository/interfaces"
	"context"
	"errors"
	"log"
)

type TicketInteractor struct {
	SearchRepository interfaces.SearchRepository
	TicketRepository interfaces.TicketRepository
}

func (t TicketInteractor) ReserveTickets(ctx context.Context, request dto.ReserveTicketRequest) (dto.ReserveTicketResponse, error) {
	var ticket dto.ReserveTicketResponse

	//Checking seats
	availableSeats, err := t.TicketRepository.CheckSeatAvailability(ctx, request.ScheduleId)
	if err != nil {
		log.Printf("TicketRepository.CheckSeatAvailability err: %v", err)
		return dto.ReserveTicketResponse{}, err
	}

	if request.NumberOfSeats > availableSeats {
		log.Printf("TicketRepository.CheckSeatAvailability err: Less seats than requested number: %v", request.NumberOfSeats)
		log.Printf("TicketRepository.CheckSeatAvailability err: available seats: %v", availableSeats)
		return dto.ReserveTicketResponse{}, errors.New("available seats not enough")
	}

	//Book seats
	ticketID, err := t.TicketRepository.ReserveTickets(ctx, request.UserId, request.NumberOfSeats, request.ScheduleId)
	if err != nil {
		log.Printf("TicketRepository.CheckSeatAvailability err: %v", err)
		return dto.ReserveTicketResponse{}, err
	}

	//Fetching details
	schedule, err := t.SearchRepository.FindScheduleByID(ctx, request.ScheduleId)
	if err != nil {
		log.Printf("TicketRepository.CheckSeatAvailability err: %v", err)
		return dto.ReserveTicketResponse{}, err
	}

	ticket.Source = schedule.Source
	ticket.Destination = schedule.Destination
	ticket.TicketId = ticketID

	return ticket, nil
}

func (t TicketInteractor) GetUsersTickets(ctx context.Context, request dto.UserTicketsRequest) ([]dto.UserTicketsResponse, error) {
	tickets, err := t.TicketRepository.GetUsersTickets(ctx, request.UserId)
	if err != nil {
		log.Printf("TicketRepository.GetUsersTickets err: %v", err)
		return nil, err
	}
	return tickets, err
}
