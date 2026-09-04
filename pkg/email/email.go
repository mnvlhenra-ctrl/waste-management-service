package email

import "context"

type Service interface {
	Send(
		ctx context.Context,
		to string,
		subject string,
		body string,
	) error
}
