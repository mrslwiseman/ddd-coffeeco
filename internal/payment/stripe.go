package payment

import (
	"errors"
	"github.com/stripe/stripe-go/v73/client"
)

type StripeService struct {
	stripeClient *client.API
}

func NewStripeService(apiKey string) (*StripeService, error) {
	if apiKey == "" {
		return nil, errors.New("invalid apiKey")
	}

	sc := &client.API{}
	sc.Init(apiKey, nil)

	return &StripeService{stripeClient: sc}, nil
}
