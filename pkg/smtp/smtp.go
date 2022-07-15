package smtp

import (
	"io"
	"net/smtp"
	"strings"
)

type Smtp struct {
	tp, host, port, user, pwd string
	c                         smtp.Client
}

func NewSmtp(typeProtocol, host, port, user, pwd string) *Smtp {
	return &Smtp{
		tp:   typeProtocol,
		host: host,
		port: port,
		user: user,
		pwd:  pwd,
	}
}

func (s *Smtp) Connect() error {
	var addr, host string
	var c *smtp.Client
	var err error

	host = strings.Join([]string{s.tp, s.host}, ".")
	addr = strings.Join([]string{host, s.port}, ":")

	if c, err = smtp.Dial(addr); err != nil {
		return err
	}

	a := smtp.PlainAuth("", s.user, s.pwd, host)

	if err = c.Auth(a); err != nil {
		return err
	}

	s.c = *c

	return nil
}

func (s *Smtp) Disconnect() error {
	return s.c.Close()
}

func (s Smtp) Send(rcpt, msg string) error {
	var w io.WriteCloser
	var err error

	if s.c.Rcpt(rcpt); err != nil {
		return err
	}

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
