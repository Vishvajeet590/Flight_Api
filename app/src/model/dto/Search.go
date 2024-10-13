package dto

import (
	"time"
)

type SearchRequest struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
}

type SearchResponse struct {
	ScheduleId            int       `json:"schedule_id"`
	Source                string    `json:"source"`
	Destination           string    `json:"destination"`
	FlightNumber          string    `json:"flight_number"`
	FlightName            string    `json:"flight_name"`
	FlightID              string    `json:"flight_id"`
	NumberOfSeats         int       `json:"number_of_seats"`
	ReservedNumberOfSeats int       `json:"reserved_number_of_seats"`
	Day                   string    `json:"day"`
	Departure             time.Time `json:"departure"`
	Arrival               time.Time `json:"arrival"`
}
