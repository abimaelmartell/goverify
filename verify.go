package main

import (
	"errors"
	"log"
	"net"
	"net/smtp"
	"strings"
	"time"
)

const (
	SMTP_PORT    = "25"
	SMTP_TIMEOUT = 30 * time.Second
)

type VerifyResult struct {
	Result       string `json:"result,omitempty"`
	MailboxExist bool   `json:"mailbox_exists"`
	IsCatchAll   bool   `json:"is_catch_all"`
	IsDisposable bool   `json:"is_disposable"`
	Email        string `json:"email"`
	Domain       string `json:"domain"`
	User         string `json:"user"`

	Client *smtp.Client `json:"-"`
}

func (self *VerifyResult) ConnectSmtp() error {
	mx, err := net.LookupMX(self.Domain)

	if err != nil {
		self.Result = "NoMxServersFound"

		return err
	}

	addr := mx[0].Host + ":" + SMTP_PORT

	conn, err := net.DialTimeout("tcp", addr, SMTP_TIMEOUT)

	if err != nil {
		self.Result = "ConnectionRefused"

		return err
	}

	client, err := smtp.NewClient(conn, mx[0].Host)

	if err != nil {
		self.Result = "NoMxServersFound"

		return err
	}

	self.Client = client

	err = self.Client.Hello("example.com")

	if err != nil {
		self.Result = "NoMxServersFound"

		return err
	}

	return nil
}

func (self *VerifyResult) ParseEmailAddress() error {
	pieces := strings.Split(self.Email, "@")

	if len(pieces) == 2 {
		self.User = pieces[0]
		self.Domain = pieces[1]

		return nil
	}

	self.Result = "InvalidEmailAddress"

	return errors.New("Invalid email address")
}

func (self *VerifyResult) CheckMailboxExist() {
	self.MailboxExist = addressExists(self.Client, self.Email)
}

func (self *VerifyResult) CheckIsCatchAll() {
	randomAddress := "n0n3x1st1ng4ddr355@" + self.Domain

	self.IsCatchAll = addressExists(self.Client, randomAddress)
}

func (self *VerifyResult) Verify() {
	var err error

	if err = self.ParseEmailAddress(); err != nil {
		return
	}

	if err = self.ConnectSmtp(); err != nil {
		log.Printf("%s\n", err.Error())

		return
	}

	self.CheckMailboxExist()

	if self.MailboxExist {
		self.CheckIsCatchAll()
	}

	self.CheckIsDisposable()
}

func (self *VerifyResult) CheckIsDisposable() {
	b, err := Asset("list.txt")

	if err != nil {
		panic(err)
	}

	s := string(b)

	self.IsDisposable = strings.Contains(s, self.Domain)
}

func addressExists(client *smtp.Client, address string) bool {
	err := client.Mail(address)

	if err != nil {
		return false
	}

	err = client.Rcpt(address)

	if err != nil {
		return false
	}

	return true
}
