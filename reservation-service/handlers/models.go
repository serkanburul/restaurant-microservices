package handlers

type createReservationRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	TableNo         int    `json:"table_no" binding:"required"`
	ReservationDate string `json:"reservation_date" binding:"required"`
	TimeSlot        string `json:"time_slot" binding:"required"`
}

type readReservationResponse struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Status          string `json:"status"`
	TableNo         int    `json:"table_no"`
	Guests          int    `json:"guests"`
	ReservationDate string `json:"reservation_date"`
	TimeSlot        string `json:"time_slot"`
	CreatedAt       string `json:"created_at"`
}

//type updateReservationRequest struct {
//	Name            string `json:"name" binding:"required"`
//	TableNo         int    `json:"table_no" binding:"required"`
//	ReservationDate string `json:"reservation_date" binding:"required"`
//	TimeSlot        string `json:"time_slot" binding:"required"`
//}
