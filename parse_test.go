package env_test

import (
	"reflect"
	"testing"

	"github.com/securehaven/env"
)

func TestParse(t *testing.T) {
	runTestCase(t, "string", "test", string("test"))
	runTestCase(t, "bool", "true", bool(true))
	runTestCase(t, "int", "1234", int(1234))
	runTestCase(t, "int8", "123", int8(123))
	runTestCase(t, "int16", "123", int16(123))
	runTestCase(t, "int32", "123", int32(123))
	runTestCase(t, "int64", "123", int64(123))
	runTestCase(t, "uint", "1234", uint(1234))
	runTestCase(t, "uint8", "123", uint8(123))
	runTestCase(t, "uint16", "123", uint16(123))
	runTestCase(t, "uint32", "123", uint32(123))
	runTestCase(t, "uint64", "123", uint64(123))
	runTestCase(t, "float32", "12.3", float32(12.3))
	runTestCase(t, "float64", "12.3", float64(12.3))
}

func runTestCase[T env.Allowed](t *testing.T, name string, value string, expected T) {
	t.Run(name, func(t *testing.T) {
		result, err := env.Parse[T](value)

		if err != nil {
			t.Error(err)
		}

		if reflect.TypeOf(expected).Kind() != reflect.TypeOf(result).Kind() {
			t.Errorf(
				"wrong return type: expected=%s, received=%s",
				reflect.TypeOf(expected).Name(),
				reflect.TypeOf(result).Name(),
			)
		}

		if expected != result {
			t.Errorf(
				"wrong result: expected=%v, received=%v",
				expected,
				result,
			)
		}
	})
}
