CREATE TABLE IF NOT EXISTS public.users
(
    id    BIGSERIAL PRIMARY KEY,
    email VARCHAR(250),
    name  varchar(255)
);


CREATE TABLE IF NOT EXISTS public.flight
(
    flight_id       BIGSERIAL PRIMARY KEY,
    flight_number   VARCHAR(250),
    flight_name     varchar(255),
    number_of_seats int
);


CREATE TABLE IF NOT EXISTS public.airport
(
    airport_id BIGSERIAL PRIMARY KEY,
    airport_city varchar(255),
    airport_code varchar(255),
    airport_name varchar(255)
);

CREATE TABLE IF NOT EXISTS public.Schedule(
    schedule_id BIGSERIAL PRIMARY KEY,
    source_airport_code varchar(255),
    destination_airport_code varchar(255),
    arrival_time timestamp without time zone,
    departure_time timestamp without time zone,
    day_week int,
    flight_id int,
    reserved_count int,
    constraint flight_fk FOREIGN KEY (flight_id) references flight(flight_id)
);


CREATE TABLE IF NOT EXISTS public.tickets(
    ticket_id text PRIMARY KEY,
    user_id int,
    schedule_id int,
    reserved_seats int
);


