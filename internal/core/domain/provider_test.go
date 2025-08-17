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
				WithProviderID(uuid.New()),
				WithProviderName(faker.Username()),
				WithProviderAccountID(uuid.New()),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(faker.Password()),
			},
		},
		{
			name: "valid_provider_with_alphanum_name",
			opts: []ProviderOpt{
				WithProviderID(uuid.New()),
				WithProviderName("test123"),
				WithProviderAccountID(uuid.New()),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(faker.Password()),
			},
		},
		{
			name: "valid_provider_with_single_char_name",
			opts: []ProviderOpt{
				WithProviderID(uuid.New()),
				WithProviderName("a"),
				WithProviderAccountID(uuid.New()),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(faker.Password()),
			},
		},
		{
			name: "valid_provider_with_max_length_name",
			opts: []ProviderOpt{
				WithProviderID(uuid.New()),
				WithProviderName(strings.Repeat("a", 50)),
				WithProviderAccountID(uuid.New()),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(faker.Password()),
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
				WithProviderName(faker.Username()),
				WithProviderAccountID(validAccountID),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(validApiKey),
			},
		},
		{
			name: "missing_name",
			opts: []ProviderOpt{
				WithProviderID(validID),
				WithProviderAccountID(validAccountID),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(validApiKey),
			},
		},
		{
			name: "empty_name",
			opts: []ProviderOpt{
				WithProviderID(validID),
				WithProviderName(""),
				WithProviderAccountID(validAccountID),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(validApiKey),
			},
		},
		{
			name: "name_with_special_chars",
			opts: []ProviderOpt{
				WithProviderID(validID),
				WithProviderName("test@provider"),
				WithProviderAccountID(validAccountID),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(validApiKey),
			},
		},
		{
			name: "name_too_long",
			opts: []ProviderOpt{
				WithProviderID(validID),
				WithProviderName(strings.Repeat("a", 51)),
				WithProviderAccountID(validAccountID),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(validApiKey),
			},
		},
		{
			name: "missing_account_id",
			opts: []ProviderOpt{
				WithProviderID(validID),
				WithProviderName(faker.Username()),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(validApiKey),
			},
		},
		{
			name: "missing_type",
			opts: []ProviderOpt{
				WithProviderID(validID),
				WithProviderName(faker.Username()),
				WithProviderAccountID(validAccountID),
				WithProviderApiKey(validApiKey),
			},
		},
		{
			name: "invalid_type",
			opts: []ProviderOpt{
				WithProviderID(validID),
				WithProviderName(faker.Username()),
				WithProviderAccountID(validAccountID),
				WithProviderType(ProviderType("invalid_provider")),
				WithProviderApiKey(validApiKey),
			},
		},
		{
			name: "missing_api_key",
			opts: []ProviderOpt{
				WithProviderID(validID),
				WithProviderName(faker.Username()),
				WithProviderAccountID(validAccountID),
				WithProviderType(ProviderTypeOpenRouter),
			},
		},
		{
			name: "empty_api_key",
			opts: []ProviderOpt{
				WithProviderID(validID),
				WithProviderName(faker.Username()),
				WithProviderAccountID(validAccountID),
				WithProviderType(ProviderTypeOpenRouter),
				WithProviderApiKey(""),
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

	t.Run("WithProviderID", func(t *testing.T) {
		t.Parallel()

		id := uuid.New()
		provider := &Provider{}
		
		opt := WithProviderID(id)
		opt(provider)

		assert.Equal(t, id, provider.ID)
	})

	t.Run("WithProviderName", func(t *testing.T) {
		t.Parallel()

		name := faker.Username()
		provider := &Provider{}
		
		opt := WithProviderName(name)
		opt(provider)

		assert.Equal(t, name, provider.Name)
	})

	t.Run("WithProviderAccountID", func(t *testing.T) {
		t.Parallel()

		accountID := uuid.New()
		provider := &Provider{}
		
		opt := WithProviderAccountID(accountID)
		opt(provider)

		assert.Equal(t, accountID, provider.AccountID)
	})

	t.Run("WithProviderType", func(t *testing.T) {
		t.Parallel()

		provider := &Provider{}
		
		opt := WithProviderType(ProviderTypeOpenRouter)
		opt(provider)

		assert.Equal(t, ProviderTypeOpenRouter, provider.Type)
	})

	t.Run("WithProviderApiKey", func(t *testing.T) {
		t.Parallel()

		apiKey := faker.Password()
		provider := &Provider{}
		
		opt := WithProviderApiKey(apiKey)
		opt(provider)

		assert.Equal(t, apiKey, provider.ApiKey)
	})
}

func TestProviderValidationEdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("name_with_numbers_only", func(t *testing.T) {
		t.Parallel()

		provider, err := NewProvider(
			WithProviderID(uuid.New()),
			WithProviderName("123456"),
			WithProviderAccountID(uuid.New()),
			WithProviderType(ProviderTypeOpenRouter),
			WithProviderApiKey(faker.Password()),
		)

		assert.NoError(t, err)
		assert.NotNil(t, provider)
	})

	t.Run("name_with_mixed_alphanumeric", func(t *testing.T) {
		t.Parallel()

		provider, err := NewProvider(
			WithProviderID(uuid.New()),
			WithProviderName("test123provider"),
			WithProviderAccountID(uuid.New()),
			WithProviderType(ProviderTypeOpenRouter),
			WithProviderApiKey(faker.Password()),
		)

		assert.NoError(t, err)
		assert.NotNil(t, provider)
	})

	t.Run("name_with_spaces", func(t *testing.T) {
		t.Parallel()

		provider, err := NewProvider(
			WithProviderID(uuid.New()),
			WithProviderName("test provider"),
			WithProviderAccountID(uuid.New()),
			WithProviderType(ProviderTypeOpenRouter),
			WithProviderApiKey(faker.Password()),
		)

		assert.Error(t, err)
		assert.Nil(t, provider)
	})
}

func TestProviderTypeConstant(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "open_router", string(ProviderTypeOpenRouter))
}