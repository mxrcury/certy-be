package mail

import "errors"

type SendEmailInput struct {
	To      string
	Subject string
	Content string
}

func (i *SendEmailInput) Validate() error {
	if i.To == "" {
		return errors.New("empty to")
	}

	if i.Subject == "" {
		return errors.New("empty subject")
	}

	if i.Content == "" {
		return errors.New("empty content")
	}

	return nil
}
