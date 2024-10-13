package repository

import (
	"Flight_Api/app/database"
	"Flight_Api/app/src/model/dto"
	"context"
	"database/sql"
	"fmt"
)

type SearchRepository struct{}

func (s SearchRepository) FindFlightByLocation(ctx context.Context, source, destination, date string) ([]dto.SearchResponse, error) {
	db := database.DB
	var result []dto.SearchResponse

	query := fmt.Sprintf(`SELECT schedule_id,source_airport_code, destination_airport_code, arrival_time, departure_time, reserved_count, flight_number, flight_name, flight.flight_id, number_of_seats FROM schedule
    								INNER JOIN flight ON public.schedule.flight_id = public.flight.flight_id
         							WHERE source_airport_code = $1 AND destination_airport_code = $2 AND DATE(departure_time) = date($3);`)

	rows, err := db.QueryContext(ctx, query, source, destination, date)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var flighInfo dto.SearchResponse
		rows.Scan(
			&flighInfo.ScheduleId,
			&flighInfo.Source,
			&flighInfo.Destination,
			&flighInfo.Arrival,
			&flighInfo.Departure,
			&flighInfo.ReservedNumberOfSeats,
			&flighInfo.FlightNumber,
			&flighInfo.FlightName,
			&flighInfo.FlightID,
			&flighInfo.NumberOfSeats)

		flighInfo.Day = flighInfo.Departure.Weekday().String()
		result = append(result, flighInfo)
	}
	return result, nil
}

func (s SearchRepository) FindScheduleByID(ctx context.Context, scheduleID int) (dto.SearchResponse, error) {
	db := database.DB
	var flightInfo dto.SearchResponse

	query := fmt.Sprintf(`SELECT schedule_id,source_airport_code, destination_airport_code, arrival_time, departure_time, reserved_count, flight_number, flight_name, flight.flight_id, number_of_seats FROM schedule
    								INNER JOIN flight ON public.schedule.flight_id = public.flight.flight_id
         							WHERE schedule_id = $1`)

	err := db.QueryRowContext(ctx, query, scheduleID).Scan(
		&flightInfo.ScheduleId,
		&flightInfo.Source,
		&flightInfo.Destination,
		&flightInfo.Arrival,
		&flightInfo.Departure,
		&flightInfo.ReservedNumberOfSeats,
		&flightInfo.FlightNumber,
		&flightInfo.FlightName,
		&flightInfo.FlightID,
		&flightInfo.NumberOfSeats,
	)
	flightInfo.Day = flightInfo.Departure.Weekday().String()

	if err != nil {
		if err == sql.ErrNoRows {
			return flightInfo, nil
		}
		return dto.SearchResponse{}, err
	}

	return flightInfo, nil
}
