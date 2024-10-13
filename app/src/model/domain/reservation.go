package domain

import "time"

type Schedule struct {
	ScheduleId         int
	SourceAirport      int
	DestinationAirport int
	ArrivalTime        time.Time
	DepartureTime      time.Time
	DayOfWeek          int
	FlightID           int
	ReservedSeats      int
}

type Ticket struct {
	TicketId      int
	UserId        int
	ScheduleId    int
	ReservedSeats int
}
