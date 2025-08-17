package domain

import (
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProviderIfValidInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		opts []ProviderOpt
	}{
		{
			name: "valid_provider_with_all_fields",
			opts: []ProviderOpt{
				WithID(uuid.New()),
				WithName(faker.Username()),
				WithAccountID(uuid.New()),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(faker.Password()),
			},
		},
		{
			name: "valid_provider_with_alphanum_name",
			opts: []ProviderOpt{
				WithID(uuid.New()),
				WithName("test123"),
				WithAccountID(uuid.New()),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(faker.Password()),
			},
		},
		{
			name: "valid_provider_with_single_char_name",
			opts: []ProviderOpt{
				WithID(uuid.New()),
				WithName("a"),
				WithAccountID(uuid.New()),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(faker.Password()),
			},
		},
		{
			name: "valid_provider_with_max_length_name",
			opts: []ProviderOpt{
				WithID(uuid.New()),
				WithName(strings.Repeat("a", 50)),
				WithAccountID(uuid.New()),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(faker.Password()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			provider, err := NewProvider(tt.opts...)

			assert.NoError(t, err)
			require.NotNil(t, provider)
			assert.NotEqual(t, uuid.Nil, provider.ID)
			assert.NotEmpty(t, provider.Name)
			assert.NotEqual(t, uuid.Nil, provider.AccountID)
			assert.Equal(t, ProviderTypeOpenRouter, provider.Type)
			assert.NotEmpty(t, provider.ApiKey)
		})
	}
}

func TestNewProviderIfInvalidInput(t *testing.T) {
	t.Parallel()

	validID := uuid.New()
	validAccountID := uuid.New()
	validApiKey := faker.Password()

	tests := []struct {
		name string
		opts []ProviderOpt
	}{
		{
			name: "missing_id",
			opts: []ProviderOpt{
				WithName(faker.Username()),
				WithAccountID(validAccountID),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(validApiKey),
			},
		},
		{
			name: "missing_name",
			opts: []ProviderOpt{
				WithID(validID),
				WithAccountID(validAccountID),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(validApiKey),
			},
		},
		{
			name: "empty_name",
			opts: []ProviderOpt{
				WithID(validID),
				WithName(""),
				WithAccountID(validAccountID),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(validApiKey),
			},
		},
		{
			name: "name_with_special_chars",
			opts: []ProviderOpt{
				WithID(validID),
				WithName("test@provider"),
				WithAccountID(validAccountID),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(validApiKey),
			},
		},
		{
			name: "name_too_long",
			opts: []ProviderOpt{
				WithID(validID),
				WithName(strings.Repeat("a", 51)),
				WithAccountID(validAccountID),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(validApiKey),
			},
		},
		{
			name: "missing_account_id",
			opts: []ProviderOpt{
				WithID(validID),
				WithName(faker.Username()),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(validApiKey),
			},
		},
		{
			name: "missing_type",
			opts: []ProviderOpt{
				WithID(validID),
				WithName(faker.Username()),
				WithAccountID(validAccountID),
				WithApiKey(validApiKey),
			},
		},
		{
			name: "invalid_type",
			opts: []ProviderOpt{
				WithID(validID),
				WithName(faker.Username()),
				WithAccountID(validAccountID),
				WithType(ProviderType("invalid_provider")),
				WithApiKey(validApiKey),
			},
		},
		{
			name: "missing_api_key",
			opts: []ProviderOpt{
				WithID(validID),
				WithName(faker.Username()),
				WithAccountID(validAccountID),
				WithType(ProviderTypeOpenRouter),
			},
		},
		{
			name: "empty_api_key",
			opts: []ProviderOpt{
				WithID(validID),
				WithName(faker.Username()),
				WithAccountID(validAccountID),
				WithType(ProviderTypeOpenRouter),
				WithApiKey(""),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			provider, err := NewProvider(tt.opts...)

			assert.Error(t, err)
			assert.Nil(t, provider)
		})
	}
}

func TestProviderOptFunctions(t *testing.T) {
	t.Parallel()

	t.Run("WithID", func(t *testing.T) {
		t.Parallel()

		id := uuid.New()
		provider := &Provider{}
		
		opt := WithID(id)
		opt(provider)

		assert.Equal(t, id, provider.ID)
	})

	t.Run("WithName", func(t *testing.T) {
		t.Parallel()

		name := faker.Username()
		provider := &Provider{}
		
		opt := WithName(name)
		opt(provider)

		assert.Equal(t, name, provider.Name)
	})

	t.Run("WithAccountID", func(t *testing.T) {
		t.Parallel()

		accountID := uuid.New()
		provider := &Provider{}
		
		opt := WithAccountID(accountID)
		opt(provider)

		assert.Equal(t, accountID, provider.AccountID)
	})

	t.Run("WithType", func(t *testing.T) {
		t.Parallel()

		provider := &Provider{}
		
		opt := WithType(ProviderTypeOpenRouter)
		opt(provider)

		assert.Equal(t, ProviderTypeOpenRouter, provider.Type)
	})

	t.Run("WithApiKey", func(t *testing.T) {
		t.Parallel()

		apiKey := faker.Password()
		provider := &Provider{}
		
		opt := WithApiKey(apiKey)
		opt(provider)

		assert.Equal(t, apiKey, provider.ApiKey)
	})
}

func TestProviderValidationEdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("name_with_numbers_only", func(t *testing.T) {
		t.Parallel()

		provider, err := NewProvider(
			WithID(uuid.New()),
			WithName("123456"),
			WithAccountID(uuid.New()),
			WithType(ProviderTypeOpenRouter),
			WithApiKey(faker.Password()),
		)

		assert.NoError(t, err)
		assert.NotNil(t, provider)
	})

	t.Run("name_with_mixed_alphanumeric", func(t *testing.T) {
		t.Parallel()

		provider, err := NewProvider(
			WithID(uuid.New()),
			WithName("test123provider"),
			WithAccountID(uuid.New()),
			WithType(ProviderTypeOpenRouter),
			WithApiKey(faker.Password()),
		)

		assert.NoError(t, err)
		assert.NotNil(t, provider)
	})

	t.Run("name_with_spaces", func(t *testing.T) {
		t.Parallel()

		provider, err := NewProvider(
			WithID(uuid.New()),
			WithName("test provider"),
			WithAccountID(uuid.New()),
			WithType(ProviderTypeOpenRouter),
			WithApiKey(faker.Password()),
		)

		assert.Error(t, err)
		assert.Nil(t, provider)
	})
}

func TestProviderTypeConstant(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "open_router", string(ProviderTypeOpenRouter))
}