package env_test

import (
	"errors"
	"testing"

	"github.com/securehaven/env"
)

func TestGet(t *testing.T) {
	cases := []struct {
		name          string
		env           string
		value         string
		fallback      string
		expectedEnv   string
		expectedValue string
	}{
		{
			name:          "nominal",
			env:           "TEST",
			value:         "value",
			fallback:      "fallback",
			expectedEnv:   "TEST",
			expectedValue: "value",
		},
		{
			name:          "fallback",
			env:           "TEST",
			value:         "value",
			fallback:      "fallback",
			expectedEnv:   "NOT_TEST",
			expectedValue: "fallback",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Setenv(c.env, c.value)

			result := env.Get(c.expectedEnv, c.fallback)

			if result != c.expectedValue {
				t.Errorf(
					"wrong result: expected=%v, received=%v",
					c.expectedValue,
					result,
				)
			}
		})
	}
}

func TestGetStrict(t *testing.T) {
	cases := []struct {
		name          string
		env           string
		value         string
		fallback      string
		expectedEnv   string
		expectedValue string
		expectedErr   error
	}{
		{
			name:          "nominal",
			env:           "TEST",
			value:         "value",
			fallback:      "fallback",
			expectedEnv:   "TEST",
			expectedValue: "value",
			expectedErr:   nil,
		},
		{
			name:          "fallback",
			env:           "TEST",
			value:         "value",
			fallback:      "fallback",
			expectedEnv:   "NOT_TEST",
			expectedValue: "fallback",
			expectedErr:   env.ErrNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Setenv(c.env, c.value)

			result, err := env.GetStrict(c.expectedEnv, c.fallback)

			if !errors.Is(err, c.expectedErr) {
				t.Error(err)
			}

			if result != c.expectedValue {
				t.Errorf(
					"wrong result: expected=%v, received=%v",
					c.expectedValue,
					result,
				)
			}
		})
	}
}
