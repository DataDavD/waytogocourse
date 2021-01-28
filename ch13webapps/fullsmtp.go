package main

import (
	"fmt"
	"net/smtp"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// serverName URI to smtp server
func (s *smtpServer) serverName() string {
	return s.host + ":" + s.port
}

func main() {
	// Sender data.
	from := "sender email address"
	password := "password" // you can enter original password or password generated with App password
	// Receiver email address.
	to := []string{
		"first receiver email address",
		//"second receiver email address",
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	message := []byte("Enter the message you want to send.")
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
