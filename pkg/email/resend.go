package email

import (
	"context"
	"fmt"
	"os"

	"github.com/resend/resend-go/v2"
)

type ResendService struct {
	client *resend.Client
	from   string
}

func NewResendService() *ResendService {

	apiKey := os.Getenv("RESEND_API_KEY")

	if apiKey == "" {
		panic("RESEND_API_KEY is not set")
	}

	from := os.Getenv("RESEND_FROM_EMAIL")

	if from == "" {
		from = "onboarding@resend.dev"
	}

	return &ResendService{
		client: resend.NewClient(apiKey),
		from:   from,
	}
}

func (s *ResendService) Send(
	ctx context.Context,
	to string,
	subject string,
	body string,
) error {

	params := &resend.SendEmailRequest{
		From:    s.from,
		To:      []string{to},
		Subject: subject,
		Html:    body,
	}

	_, err := s.client.Emails.SendWithContext(
		ctx,
		params,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to send email: %w",
			err,
		)
	}

	return nil
}
