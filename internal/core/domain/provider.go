package domain

import (
	"flow-run/internal/lib/validator"

	"github.com/google/uuid"
)

type ProviderType string

const (
	ProviderTypeOpenRouter = "open_router"
)

type Provider struct {
	ID        uuid.UUID    `json:"id" validate:"required"`
	Name      string       `json:"name" validate:"required,alphanum,min=1,max=50"`
	AccountID uuid.UUID    `json:"account_id" validate:"required"`
	Type      ProviderType `json:"type" validate:"oneof=open_router"`
	ApiKey    string       `json:"api_key" validate:"required"`
}

type ProviderOpt func(*Provider)

func WithID(id uuid.UUID) ProviderOpt {
	return func(p *Provider) {
		p.ID = id
	}
}

func WithName(name string) ProviderOpt {
	return func(p *Provider) {
		p.Name = name
	}
}

func WithAccountID(accountID uuid.UUID) ProviderOpt {
	return func(p *Provider) {
		p.AccountID = accountID
	}
}

func WithType(providerType ProviderType) ProviderOpt {
	return func(p *Provider) {
		p.Type = providerType
	}
}

func WithApiKey(apiKey string) ProviderOpt {
	return func(p *Provider) {
		p.ApiKey = apiKey
	}
}

func NewProvider(opts ...ProviderOpt) (*Provider, error) {

	p := &Provider{}
	for _, opt := range opts {
		opt(p)
	}

	return validator.Struct(p)
}
