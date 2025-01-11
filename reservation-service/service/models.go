package service

import (
	"database/sql"
	"time"
)

type Payload struct {
	Token           string
	Name            string
	Email           string
	Status          string
	TableNo         int32
	TimeSlotID      int32
	ReservationDate time.Time
}

type CreatePayload struct {
	Name            string
	Email           string
	TableNo         int
	ReservationDate time.Time
	TimeSlot        string
}

type UpdatePayload struct {
	Token           string
	Name            string
	TableNo         int
	ReservationDate string
	TimeSlot        string
}

type Reservation struct {
	Name            string
	Email           string
	Status          string
	TableNo         sql.NullInt32
	Capacity        int32
	TimeSlot        string
	ReservationDate time.Time
	CreatedAt       time.Time
}

type ListReservations struct {
	TableNo   int
	StartTime string
}
