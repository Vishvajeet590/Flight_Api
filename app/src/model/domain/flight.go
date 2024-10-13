package domain

type Flight struct {
	FlightID      int    `json:"flight_id"`
	FlightNumber  string `json:"flight_number"`
	FlightName    string `json:"flight_name"`
	NumberOfSeats int    `json:"number_of_seats"`
}
