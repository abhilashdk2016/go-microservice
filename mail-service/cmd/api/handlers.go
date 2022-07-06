package main

import (
	"log"
	"net/http"
)

type MailMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Send Mail")
	var requestPayload MailMessage
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("Read JSON")
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Data:    requestPayload.Message,
		Subject: requestPayload.Subject,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		log.Println("Send SMTP Message")
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Mail Sent to" + requestPayload.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
