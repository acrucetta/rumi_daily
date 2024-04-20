package main

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

func sendEmail(poem string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed loading .env")
	}
	pass := os.Getenv("APP_PASSWORD")
	auth := LoginAuth("acrucettanieto", pass)

	to := []string{"andres.crucetta@hey.com"}
	msg := []byte("To: andres.crucetta@hey.com\r\n" +
		"Subject: Daily poem\r\n" +
		"\r\n" + poem)
	err = smtp.SendMail("smtp.gmail.com:587", auth, "acrucettanieto@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
