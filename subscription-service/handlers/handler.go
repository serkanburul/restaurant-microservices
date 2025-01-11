package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
	"subscription-service/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type createSubscriptionRequest struct {
	MailAddress string `json:"mailAddress"`
}

func (h *Handler) SendMail(c echo.Context) error {
	subject := c.FormValue("subject")
	body := c.FormValue("body")
	err := h.service.SendMail(subject, body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("an error occurred"))
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) CreateSubscription(c echo.Context) error {
	c.Request().Header.Set("Content-Type", "application/json")
	var req createSubscriptionRequest
	err := c.Bind(&req)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("something went wrong"))
	}
	if req.MailAddress != "" {
		err = h.service.CreateSubscription(req.MailAddress)
		if err != nil {
			log.Println(err)
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusConflict, "this email address already exists")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, errors.New("something went wrong"))
		}

		return c.JSON(http.StatusCreated, "Subscribed successfully.")
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("empty"))
	}
}
