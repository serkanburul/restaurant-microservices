package main

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"reservation-service/db"
	"reservation-service/handlers"
	"reservation-service/service"
)

func main() {
	pConn, err := DBConn()
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.New(
		"file://migrations",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	newDB := db.NewDB(pConn)
	newService := service.NewService(newDB)
	newHandler := handlers.NewHandler(newService)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// CRUD FOR SINGLE RESERVATION
	e.POST("/reservation", newHandler.CreateReservation)
	e.GET("/reservation/:token", newHandler.ReadReservation)
	//e.PATCH("/reservation/:token", newHandler.UpdateReservation)
	e.DELETE("/reservation/:token", newHandler.DeleteReservation)

	// GET BY DATE AND CAPACITY
	e.GET("/reservation/:date/:capacity", newHandler.GetReservationsByDateAndCapacity)
	
	// GET SLOTS
	e.GET("/reservation/slots", newHandler.GetSlots)

	e.Logger.Fatal(e.Start(":1323"))
}

func DBConn() (*pgxpool.Pool, error) {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	dbPool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	return dbPool, nil
}
