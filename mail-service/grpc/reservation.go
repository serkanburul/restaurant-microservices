package grpc_server

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"gopkg.in/gomail.v2"
	"log"
	"mail-service/proto"
	"text/template"
)

type ReservationMailTemplateParams struct {
	Token           string
	Name            string
	ReservationDate string
	TimeSlot        string
}

//go:embed templates/*
var templateFS embed.FS

func (m *MailServer) SendReservationMail(ctx context.Context, req *proto.SendReservationMailRequest) (*proto.SendMailResponse, error) {
	var emailPath string
	var subject string

	switch req.ReservationMailEntry.Type {
	case "CREATE":
		emailPath = "templates/reservation/create_email_template.html"
		subject = "RESERVATION CREATION"
	case "UPDATE":
		emailPath = "templates/reservation/update_email_template.html"
		subject = "UPDATE RESERVATION"
	case "DELETE":
		emailPath = "templates/reservation/delete_email_template.html"
		subject = "DELETE RESERVATION"
	}

	t, err := template.ParseFS(templateFS, emailPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tmail := ReservationMailTemplateParams{
		Token:           req.ReservationMailEntry.Token,
		Name:            req.ReservationMailEntry.Name,
		ReservationDate: req.ReservationMailEntry.ReservationDate,
		TimeSlot:        req.ReservationMailEntry.TimeSlot,
	}

	var body bytes.Buffer
	err = t.Execute(&body, tmail)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	SMTPSettings, err := SMTPConfig()
	if err != nil {
		return nil, errors.New("could not get SMTP settings: " + err.Error())
	}

	log.Println("SMTP Settings: ", SMTPSettings)
	log.Println("req: ", req.ReservationMailEntry)

	mail := gomail.NewMessage()
	mail.SetHeader("From", SMTPSettings.FROM)
	mail.SetHeader("To", req.ReservationMailEntry.To)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body.String())

	d := gomail.NewDialer(SMTPSettings.HOST, SMTPSettings.PORT, SMTPSettings.USER, SMTPSettings.PASS)

	if err := d.DialAndSend(mail); err != nil {
		log.Printf("Error sending mail: %v", err)
		return &proto.SendMailResponse{Error: true, Message: "Error sending mail"}, err
	}

	log.Println("Mail sent successfully")

	return &proto.SendMailResponse{Error: false, Message: "Mail sent successfully"}, nil

}
