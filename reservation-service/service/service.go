package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"reservation-service/db"
	"reservation-service/infrastructure/mail"
	"time"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetSlots() ([]string, error) {
	ctx := context.Background()
	slots, err  := s.db.ListTimeSlots(ctx)
	if err != nil {
		return nil, err
	}
	formattedSlots := make([]string, len(slots))
	for i, slot := range slots {
		formattedSlots[i] = slot.Format("15:04:05")
		if err != nil {
			return nil, err
		}
	}
	return formattedSlots, nil
}

func (s *Service) CreateReservation(p CreatePayload) error {
	ctx := context.Background()
	parsedTime, err := time.Parse("15:04:05", p.TimeSlot)
	if err != nil {
		return err
	}

	slot, err := s.db.GetTimeSlot(ctx, parsedTime)
	if err != nil {
		log.Println("Error getting time slot: ", err.Error())
		return errors.New("error getting time slot by id")
	}

	reservation, err := s.db.CreateTableReservation(ctx, db.CreateTableReservationParams{
		Token:           GenerateSecureUniqueCode(),
		Name:            p.Name,
		Email:           p.Email,
		Status:          "PENDING",
		TableNo:         sql.NullInt32{Int32: int32(p.TableNo), Valid: true},
		ReservationDate: p.ReservationDate,
		TimeSlotID:      sql.NullInt32{Int32: slot.ID, Valid: true},
	})
	if err != nil {
		log.Println("Error creating reservation table: ", err.Error())
		return errors.New("Error creating reservation")
	}

	err = mail.SendMail(mail.Mail{
		Type:            "CREATE",
		Token:           reservation.Token,
		Name:            p.Name,
		Email:           p.Email,
		Capacity:        p.TableNo,
		ReservationDate: p.ReservationDate.Format("2006-01-02"),
		TimeSlot:        p.TimeSlot,
	})
	if err != nil {
		return errors.New("error sending mail")
	}

	return nil

}

func (s *Service) GetReservationsByDateAndCapacity(date string, capacity int32) ([]ListReservations, error) {
	ctx := context.Background()
	var res []ListReservations

	dater, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println("Error parsing date: ", err.Error())
		return res, errors.New("error parsing date")
	}

	reservations, err := s.db.GetReservationsByDateAndCapacity(ctx, db.GetReservationsByDateAndCapacityParams{
		ReservationDate: dater,
		Capacity:        capacity,
	})
	if err != nil {
		log.Println("Error getting reservations: ", err.Error())
		return nil, err
	}

	for _, reservation := range reservations {
		t := reservation.TableNo.Int32
		res = append(res, ListReservations{
			TableNo:   int(t),
			StartTime: reservation.StartTime.Format("15:04:05"),
		})
	}
	return res, nil
}

func (s *Service) GetReservationByToken(token string) (Reservation, error) {
	ctx := context.Background()

	reservation, err := s.db.GetReservationByToken(ctx, token)
	if err != nil {
		return Reservation{}, err
	}
	return Reservation{
		Name:            reservation.Name,
		Email:           reservation.Email,
		Status:          reservation.Status,
		TableNo:         reservation.TableNo,
		Capacity:        reservation.Capacity,
		ReservationDate: reservation.ReservationDate,
		TimeSlot:        reservation.StartTime.Format("15:04"),
		CreatedAt:       reservation.CreatedAt.Time,
	}, nil
}

//func (s *Service) UpdateReservation(payload UpdatePayload) (Reservation, error) {
//	ctx := context.Background()
//
//	sqlTableNo := sql.NullInt32{
//		Int32: int32(payload.TableNo),
//		Valid: true,
//	}
//
//	reservationDate, err := time.Parse("2006-01-02", payload.ReservationDate)
//	if err != nil {
//		return Reservation{}, err
//	}
//
//	parse, err := time.Parse("2006-01-02", payload.TimeSlot)
//	if err != nil {
//		return Reservation{}, err
//	}
//
//	timeSlot, err := s.db.GetTimeSlot(ctx, parse)
//
//	sqlTimeSlotId := sql.NullInt32{
//		Int32: timeSlot.ID,
//		Valid: true,
//	}
//
//	reservation, err := s.db.UpdateReservation(ctx, db.UpdateReservationParams{
//		Token:           payload.Token,
//		Name:            payload.Name,
//		Status:          "PENDING",
//		TableNo:         sqlTableNo,
//		ReservationDate: reservationDate,
//		TimeSlotID:      sqlTimeSlotId,
//	})
//	if err != nil {
//		return Reservation{}, err
//	}
//
//	tableByID, err := s.db.GetTableByID(ctx, sqlTableNo.Int32)
//	if err != nil {
//		return Reservation{}, err
//	}
//
//	err = mail.SendMail(mail.Mail{
//		Token:           reservation.Token,
//		Type:            "UPDATE",
//		Name:            reservation.Name,
//		Email:           reservation.Email,
//		Capacity:        int(tableByID.Capacity),
//		ReservationDate: reservation.ReservationDate.Format("2006-01-02"),
//		TimeSlot:        payload.TimeSlot,
//	})
//	if err != nil {
//		return Reservation{}, err
//	}
//
//	return Reservation{
//		Name:            reservation.Name,
//		Email:           reservation.Email,
//		Status:          reservation.Status,
//		TableNo:         sqlTableNo,
//		ReservationDate: reservation.ReservationDate,
//		CreatedAt:       reservation.CreatedAt.Time,
//	}, err
//
//}

func (s *Service) DeleteReservationByToken(token string) error {
	ctx := context.Background()

	res, err := s.db.GetReservationByToken(ctx, token)
	if err != nil {
		log.Println("Error getting reservation table: ", err.Error())
		return errors.New("error getting reservation table")
	}

	err = s.db.DeleteReservationByToken(ctx, token)
	if err != nil {
		log.Println("Error deleting reservation table: ", err.Error())
		return errors.New("error deleting reservation")
	}

	slot, err := s.db.GetTimeSlot(ctx, res.StartTime)
	if err != nil {
		log.Println("Error getting time slot: ", err.Error())
		return errors.New("error getting time slot by id")
	}

	err = mail.SendMail(mail.Mail{
		Type:            "DELETE",
		Name:            res.Name,
		Email:           res.Email,
		Capacity:        int(res.TableNo.Int32),
		ReservationDate: res.ReservationDate.Format("2006-01-02"),
		TimeSlot:        slot.StartTime.Format("15:04"),
	})
	if err != nil {
		return errors.New("error sending mail")
	}

	return nil
}
