package smtp

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
)

type Smtp struct {
	host, port, user, pwd string
	c                     smtp.Client
}

func NewSmtp(servername, user, pwd string) *Smtp {

	host, port, _ := net.SplitHostPort(servername)

	return &Smtp{
		host: host,
		port: port,
		user: user,
		pwd:  pwd,
	}
}

func (s *Smtp) Connect() error {
	var addr string
	var conn *tls.Conn
	var c *smtp.Client
	var err error

	addr = strings.Join([]string{s.host, s.port}, ":")
	auth := smtp.PlainAuth("", s.user, s.pwd, s.host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.host,
	}

	if conn, err = tls.Dial("tcp", addr, tlsconfig); err != nil {
		return err
	}

	if c, err = smtp.NewClient(conn, s.host); err != nil {
		return err
	}

	if err = c.Auth(auth); err != nil {
		return err
	}

	s.c = *c

	return nil
}

func (s *Smtp) Disconnect() error {
	return s.c.Close()
}

func (s *Smtp) Send(rcpt, msg, subj string) error {
	var err error

	if err = s.setSenderAndReceiver(rcpt); err != nil {
		return err
	}

	if err = s.writeMsg(s.makeBody(rcpt, msg, subj)); err != nil {
		return err
	}

	return nil
}

func (s *Smtp) setSenderAndReceiver(rcpt string) error {
	var to *mail.Address
	var err error

	if to, err = mail.ParseAddress(rcpt); err != nil {
		return err
	}

	if err = s.c.Mail(s.user); err != nil {
		return err
	}

	if err = s.c.Rcpt(to.Address); err != nil {
		return err
	}

	return nil
}

func (s *Smtp) makeBody(rcpt, msg, subj string) string {
	headers := make(map[string]string)
	headers["From"] = s.user
	headers["To"] = rcpt
	headers["Subject"] = subj

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += "\r\n" + msg

	return message
}

func (s *Smtp) writeMsg(msg string) error {
	var w io.WriteCloser
	var err error

	if w, err = s.c.Data(); err != nil {
		return err
	}

	if _, err = w.Write([]byte(msg)); err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	return nil
}
