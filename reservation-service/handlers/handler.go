package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"reservation-service/service"
	"strconv"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetSlots(c echo.Context) error {
	slots, err  := h.service.GetSlots()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, slots)
}

func (h *Handler) CreateReservation(c echo.Context) error {
	var req createReservationRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	parse, err := time.Parse("2006-01-02", req.ReservationDate)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Something went wrong")
	}

	err = h.service.CreateReservation(service.CreatePayload{
		Name:            req.Name,
		Email:           req.Email,
		TableNo:         req.TableNo,
		ReservationDate: parse,
		TimeSlot:        req.TimeSlot,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (h *Handler) ReadReservation(c echo.Context) error {
	param := c.Param("token")
	reservation, err := h.service.GetReservationByToken(param)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, readReservationResponse{
		Name:            reservation.Name,
		Email:           reservation.Email,
		Status:          reservation.Status,
		TableNo:         int(reservation.TableNo.Int32),
		Guests:          int(reservation.Capacity),
		ReservationDate: reservation.ReservationDate.Format("2006-01-02"),
		TimeSlot:        reservation.TimeSlot,
		CreatedAt:       reservation.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

//func (h *Handler) UpdateReservation(c echo.Context) error {
//	token := c.Param("token")
//	var req updateReservationRequest
//	err := c.Bind(&req)
//	if err != nil {
//		log.Println(err)
//		return echo.NewHTTPError(http.StatusBadRequest, errors.New("an error occured"))
//	}
//	reservation, err := h.service.UpdateReservation(service.UpdatePayload{
//		Token:           token,
//		Name:            req.Name,
//		TableNo:         req.TableNo,
//		ReservationDate: req.ReservationDate,
//		TimeSlot:        req.TimeSlot,
//	})
//	if err != nil {
//		log.Println(err)
//		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
//	}
//	return c.JSON(http.StatusOK, reservation)
//}

func (h *Handler) DeleteReservation(c echo.Context) error {
	param := c.Param("token")

	err := h.service.DeleteReservationByToken(param)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("an error occured"))
	}

	return c.JSON(http.StatusOK, "Deletion completed successfully")

}

func (h *Handler) GetReservationsByDateAndCapacity(c echo.Context) error {
	date := c.Param("date")
	capacity := c.Param("capacity")
	atoi, err := strconv.Atoi(capacity)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("an error occured"))
	}

	reservations, err := h.service.GetReservationsByDateAndCapacity(date, int32(atoi))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, reservations)
}
