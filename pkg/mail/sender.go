package mail

import "gopkg.in/gomail.v2"

type SMTPSender struct {
	pass     string
	port     int
	from     string
	username string
	host     string
}

func NewSMTPSender(username, pass, from, host string, port int) *SMTPSender {
	return &SMTPSender{
		pass:     pass,
		port:     port,
		from:     from,
		host:     host,
		username: username,
	}
}

func (m *SMTPSender) SendEmail(input *SendEmailInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	print("EMAIL:", input.To)

	msg := gomail.NewMessage()

	msg.SetHeader("From", m.from)
	msg.SetHeader("To", input.To)
	msg.SetHeader("Subject", input.Subject)
	msg.SetBody("text/html", input.Content)

	dialer := gomail.NewDialer(m.host, m.port, m.username, m.pass)

	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
