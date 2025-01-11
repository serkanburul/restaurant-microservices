package grpc_server

import (
	"bytes"
	"context"
	"errors"
	"gopkg.in/gomail.v2"
	"log"
	"mail-service/proto"
	"text/template"
)

type SubscriptionMailTemplateParams struct {
	Subject string
	Body    string
}

func (m *MailServer) SendSubscriptionMail(ctx context.Context, req *proto.SendSubscriptionMailRequest) (*proto.SendMailResponse, error) {
	var emailPath = "templates/subscription/subscription_template.html"
	var subject string

	t, err := template.ParseFS(templateFS, emailPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tmail := SubscriptionMailTemplateParams{
		Subject: req.SubscriptionMailEntry.Subject,
		Body:    req.SubscriptionMailEntry.Body,
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
	log.Println("req: ", req.SubscriptionMailEntry)

	mail := gomail.NewMessage()
	mail.SetHeader("From", SMTPSettings.FROM)
	mail.SetHeader("To", req.SubscriptionMailEntry.To...)
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
