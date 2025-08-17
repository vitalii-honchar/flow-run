package domain

import (
	"flow-run/internal/lib/validator"

	"github.com/google/uuid"
)

type Model struct {
	ID         uuid.UUID `json:"id" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	AccountID  uuid.UUID `json:"account_id" validate:"required"`
	ProviderID uuid.UUID `json:"provider_id" validate:"required"`
}

type ModelOpt func(*Model)

func WithModelID(id uuid.UUID) ModelOpt {
	return func(m *Model) {
		m.ID = id
	}
}

func WithModelName(name string) ModelOpt {
	return func(m *Model) {
		m.Name = name
	}
}

func WithModelAccountID(accountID uuid.UUID) ModelOpt {
	return func(m *Model) {
		m.AccountID = accountID
	}
}

func WithModelProviderID(providerID uuid.UUID) ModelOpt {
	return func(m *Model) {
		m.ProviderID = providerID
	}
}

func NewModel(opts ...ModelOpt) (*Model, error) {
	m := &Model{}
	for _, opt := range opts {
		opt(m)
	}
	return validator.Struct(m)
}
