package repository

import (
	"Flight_Api/app/database"
	"Flight_Api/app/src/model/dto"
	"Flight_Api/app/src/repository/interfaces"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type TicketRepository struct {
	SearchRepository interfaces.SearchRepository
}

func (t TicketRepository) GetMaxNumberOfSeats(ctx context.Context, flightID int) (int, error) {
	db := database.DB
	var result int

	query := fmt.Sprintf(`select flight.number_of_seats from  flight where flight_id = $1`)
	err := db.QueryRowContext(ctx, query, flightID).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		fmt.Printf("Error while fetching max number of seats: %v\n", err)
		return 0, err
	}
	return result, nil
}

func (t TicketRepository) CheckSeatAvailability(ctx context.Context, scheduleID int) (int, error) {
	maxNumberOfSeats, err := t.GetMaxNumberOfSeats(ctx, scheduleID)
	if err != nil {
		return 0, err
	}
	schedule, err := t.SearchRepository.FindScheduleByID(ctx, scheduleID)
	if err != nil {
		return 0, err
	}

	//+ve means available -ve not available
	return maxNumberOfSeats - schedule.ReservedNumberOfSeats, nil
}

func (t TicketRepository) ReserveTickets(ctx context.Context, userID, seats, scheduleID int) (string, error) {
	//Add to tickets
	db := database.DB
	ticketID := uuid.NewString()
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error while reserving tickets: %v\n", err)
		return "", err
	}
	ticketQuery := fmt.Sprintf(`INSERT INTO tickets (ticket_id, user_id, reserved_seats, schedule_id)  VALUES ('%v',$1,$2,$3);`, ticketID)
	_, err = tx.ExecContext(ctx, ticketQuery, userID, seats, scheduleID)
	if err != nil {
		log.Printf("Error while reserving tickets: %v\n", err)
		tx.Rollback()
		return "", err
	}

	//Updating the count in Schedule

	scheduleQuery := fmt.Sprintf(`UPDATE schedule SET reserved_count = reserved_count + $1 WHERE schedule_id = $2;`)
	_, err = tx.ExecContext(ctx, scheduleQuery, seats, scheduleID)
	if err != nil {
		log.Printf("Error while reserving tickets: %v\n", err)
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error while reserving tickets: %v\n", err)
		return "", err
	}

	return ticketID, nil
}

func (t TicketRepository) GetUsersTickets(ctx context.Context, userID int) ([]dto.UserTicketsResponse, error) {
	db := database.DB
	var result []dto.UserTicketsResponse

	query := fmt.Sprintf(`SELECT ticket_id,source_airport_code,destination_airport_code,arrival_time,departure_time,tickets.reserved_seats,flight_name,flight_number FROM tickets INNER JOIN public.schedule ON tickets.schedule_id = schedule.schedule_id INNER JOIN public.flight f on f.flight_id = schedule.flight_id WHERE tickets.user_id = $1;`)

	rows, err := db.QueryContext(ctx, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		log.Printf("Error while fetching tickets: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var ticket dto.UserTicketsResponse
		rows.Scan(
			&ticket.TicketId,
			&ticket.Source,
			&ticket.Destination,
			&ticket.ArrivalDate,
			&ticket.DepartureDate,
			&ticket.ReservedSeats,
			&ticket.FlightName,
			&ticket.FlightNumber,
		)

		result = append(result, ticket)
	}

	return result, nil
}
