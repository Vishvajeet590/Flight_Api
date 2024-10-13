package dto

type ReserveTicketRequest struct {
	UserId        int `json:"user_id"`
	ScheduleId    int `json:"schedule_id"`
	NumberOfSeats int `json:"number_of_seats"`
}

type ReserveTicketResponse struct {
	TicketId      string `json:"ticket_id"`
	Source        string `json:"source"`
	Destination   string `json:"destination"`
	FlightNumber  string `json:"flight_number"`
	FlightName    string `json:"flight_name"`
	DepartureDate string `json:"departure_date"`
	ArrivalDate   string `json:"arrival_date"`
	ReservedSeats int    `json:"reserved_seats"`
}

type UserTicketsRequest struct {
	UserId int `json:"user_id"`
}

type UserTicketsResponse struct {
	TicketId      string `json:"ticket_id"`
	Source        string `json:"source"`
	Destination   string `json:"destination"`
	FlightNumber  string `json:"flight_number"`
	FlightName    string `json:"flight_name"`
	DepartureDate string `json:"departure_date"`
	ArrivalDate   string `json:"arrival_date"`
	ReservedSeats int    `json:"reserved_seats"`
}
