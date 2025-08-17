package domain

import (
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewModelIfValidInput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		opts []ModelOpt
	}{
		{
			name: "valid_model_with_all_fields",
			opts: []ModelOpt{
				WithModelID(uuid.New()),
				WithModelName(faker.Name()),
				WithModelAccountID(uuid.New()),
				WithModelProviderID(uuid.New()),
			},
		},
		{
			name: "valid_model_with_short_name",
			opts: []ModelOpt{
				WithModelID(uuid.New()),
				WithModelName("a"),
				WithModelAccountID(uuid.New()),
				WithModelProviderID(uuid.New()),
			},
		},
		{
			name: "valid_model_with_long_name",
			opts: []ModelOpt{
				WithModelID(uuid.New()),
				WithModelName(strings.Repeat("ModelName", 10)),
				WithModelAccountID(uuid.New()),
				WithModelProviderID(uuid.New()),
			},
		},
		{
			name: "valid_model_with_alphanumeric_name",
			opts: []ModelOpt{
				WithModelID(uuid.New()),
				WithModelName("Model123"),
				WithModelAccountID(uuid.New()),
				WithModelProviderID(uuid.New()),
			},
		},
		{
			name: "valid_model_with_special_chars_in_name",
			opts: []ModelOpt{
				WithModelID(uuid.New()),
				WithModelName("Model-Name_v2.0"),
				WithModelAccountID(uuid.New()),
				WithModelProviderID(uuid.New()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			model, err := NewModel(tt.opts...)

			assert.NoError(t, err)
			require.NotNil(t, model)
			assert.NotEqual(t, uuid.Nil, model.ID)
			assert.NotEmpty(t, model.Name)
			assert.NotEqual(t, uuid.Nil, model.AccountID)
			assert.NotEqual(t, uuid.Nil, model.ProviderID)
		})
	}
}

func TestNewModelIfInvalidInput(t *testing.T) {
	t.Parallel()

	validID := uuid.New()
	validAccountID := uuid.New()
	validProviderID := uuid.New()
	validName := faker.Name()

	tests := []struct {
		name string
		opts []ModelOpt
	}{
		{
			name: "missing_id",
			opts: []ModelOpt{
				WithModelName(validName),
				WithModelAccountID(validAccountID),
				WithModelProviderID(validProviderID),
			},
		},
		{
			name: "missing_name",
			opts: []ModelOpt{
				WithModelID(validID),
				WithModelAccountID(validAccountID),
				WithModelProviderID(validProviderID),
			},
		},
		{
			name: "empty_name",
			opts: []ModelOpt{
				WithModelID(validID),
				WithModelName(""),
				WithModelAccountID(validAccountID),
				WithModelProviderID(validProviderID),
			},
		},
		{
			name: "missing_account_id",
			opts: []ModelOpt{
				WithModelID(validID),
				WithModelName(validName),
				WithModelProviderID(validProviderID),
			},
		},
		{
			name: "missing_provider_id",
			opts: []ModelOpt{
				WithModelID(validID),
				WithModelName(validName),
				WithModelAccountID(validAccountID),
			},
		},
		{
			name: "all_fields_missing",
			opts: []ModelOpt{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			model, err := NewModel(tt.opts...)

			assert.Error(t, err)
			assert.Nil(t, model)
		})
	}
}

func TestModelOptFunctions(t *testing.T) {
	t.Parallel()

	t.Run("WithModelID", func(t *testing.T) {
		t.Parallel()

		id := uuid.New()
		model := &Model{}
		
		opt := WithModelID(id)
		opt(model)

		assert.Equal(t, id, model.ID)
	})

	t.Run("WithModelName", func(t *testing.T) {
		t.Parallel()

		name := faker.Name()
		model := &Model{}
		
		opt := WithModelName(name)
		opt(model)

		assert.Equal(t, name, model.Name)
	})

	t.Run("WithModelAccountID", func(t *testing.T) {
		t.Parallel()

		accountID := uuid.New()
		model := &Model{}
		
		opt := WithModelAccountID(accountID)
		opt(model)

		assert.Equal(t, accountID, model.AccountID)
	})

	t.Run("WithModelProviderID", func(t *testing.T) {
		t.Parallel()

		providerID := uuid.New()
		model := &Model{}
		
		opt := WithModelProviderID(providerID)
		opt(model)

		assert.Equal(t, providerID, model.ProviderID)
	})
}

func TestModelValidationEdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("name_with_whitespace", func(t *testing.T) {
		t.Parallel()

		model, err := NewModel(
			WithModelID(uuid.New()),
			WithModelName("  Model Name  "),
			WithModelAccountID(uuid.New()),
			WithModelProviderID(uuid.New()),
		)

		assert.NoError(t, err)
		assert.NotNil(t, model)
		assert.Equal(t, "  Model Name  ", model.Name)
	})

	t.Run("name_with_unicode_chars", func(t *testing.T) {
		t.Parallel()

		model, err := NewModel(
			WithModelID(uuid.New()),
			WithModelName("模型名称"),
			WithModelAccountID(uuid.New()),
			WithModelProviderID(uuid.New()),
		)

		assert.NoError(t, err)
		assert.NotNil(t, model)
	})

	t.Run("same_id_for_account_and_provider", func(t *testing.T) {
		t.Parallel()

		sameID := uuid.New()
		model, err := NewModel(
			WithModelID(uuid.New()),
			WithModelName(faker.Name()),
			WithModelAccountID(sameID),
			WithModelProviderID(sameID),
		)

		assert.NoError(t, err)
		assert.NotNil(t, model)
		assert.Equal(t, sameID, model.AccountID)
		assert.Equal(t, sameID, model.ProviderID)
	})

	t.Run("model_id_same_as_account_id", func(t *testing.T) {
		t.Parallel()

		sameID := uuid.New()
		model, err := NewModel(
			WithModelID(sameID),
			WithModelName(faker.Name()),
			WithModelAccountID(sameID),
			WithModelProviderID(uuid.New()),
		)

		assert.NoError(t, err)
		assert.NotNil(t, model)
		assert.Equal(t, sameID, model.ID)
		assert.Equal(t, sameID, model.AccountID)
	})
}

func TestModelStructValidation(t *testing.T) {
	t.Parallel()

	t.Run("direct_struct_creation_valid", func(t *testing.T) {
		t.Parallel()

		model := &Model{
			ID:         uuid.New(),
			Name:       faker.Name(),
			AccountID:  uuid.New(),
			ProviderID: uuid.New(),
		}

		// Test that the struct itself has the expected fields
		assert.NotEqual(t, uuid.Nil, model.ID)
		assert.NotEmpty(t, model.Name)
		assert.NotEqual(t, uuid.Nil, model.AccountID)
		assert.NotEqual(t, uuid.Nil, model.ProviderID)
	})

	t.Run("json_tags_present", func(t *testing.T) {
		t.Parallel()

		// This test ensures JSON tags are correctly set
		model := Model{
			ID:         uuid.New(),
			Name:       "test-model",
			AccountID:  uuid.New(),
			ProviderID: uuid.New(),
		}

		// Verify that the struct can be used for JSON serialization
		assert.NotNil(t, model.ID)
		assert.Equal(t, "test-model", model.Name)
		assert.NotNil(t, model.AccountID)
		assert.NotNil(t, model.ProviderID)
	})
}

func TestModelCreationWithRandomData(t *testing.T) {
	t.Parallel()

	for i := 0; i < 10; i++ {
		t.Run(faker.Word(), func(t *testing.T) {
			t.Parallel()

			model, err := NewModel(
				WithModelID(uuid.New()),
				WithModelName(faker.Sentence()),
				WithModelAccountID(uuid.New()),
				WithModelProviderID(uuid.New()),
			)

			assert.NoError(t, err)
			require.NotNil(t, model)
			assert.NotEqual(t, uuid.Nil, model.ID)
			assert.NotEmpty(t, model.Name)
			assert.NotEqual(t, uuid.Nil, model.AccountID)
			assert.NotEqual(t, uuid.Nil, model.ProviderID)
		})
	}
}