package env

import (
	"errors"
	"strconv"
)

type Allowed interface {
	~string | ~bool | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

var (
	ErrFailedConversion = errors.New("can not convert value to expected type")
	ErrFailedCasting    = errors.New("can not cast value to expected type")
)

func Parse[T Allowed](value string) (T, error) {
	var (
		result T
		ok     bool
		err    error
	)

	switch any(result).(type) {
	case string:
		result, ok = any(value).(T)
	case bool:
		var v bool

		v, err = strconv.ParseBool(value)
		result, ok = any(v).(T)
	case int:
		var v int64

		v, err = strconv.ParseInt(value, 10, 0)
		result, ok = any(int(v)).(T)
	case int8:
		var v int64

		v, err = strconv.ParseInt(value, 10, 8)
		result, ok = any(int8(v)).(T)
	case int16:
		var v int64

		v, err = strconv.ParseInt(value, 10, 16)
		result, ok = any(int16(v)).(T)
	case int32:
		var v int64

		v, err = strconv.ParseInt(value, 10, 32)
		result, ok = any(int32(v)).(T)
	case int64:
		var v int64

		v, err = strconv.ParseInt(value, 10, 64)
		result, ok = any(v).(T)
	case uint:
		var v uint64

		v, err = strconv.ParseUint(value, 10, 0)
		result, ok = any(uint(v)).(T)
	case uint8:
		var v uint64

		v, err = strconv.ParseUint(value, 10, 8)
		result, ok = any(uint8(v)).(T)
	case uint16:
		var v uint64

		v, err = strconv.ParseUint(value, 10, 16)
		result, ok = any(uint16(v)).(T)
	case uint32:
		var v uint64

		v, err = strconv.ParseUint(value, 10, 32)
		result, ok = any(uint32(v)).(T)
	case uint64:
		var v uint64

		v, err = strconv.ParseUint(value, 10, 64)
		result, ok = any(v).(T)
	case float32:
		var v float64

		v, err = strconv.ParseFloat(value, 32)
		result, ok = any(float32(v)).(T)
	case float64:
		var v float64

		v, err = strconv.ParseFloat(value, 64)
		result, ok = any(v).(T)
	default:
		err = ErrFailedConversion
	}

	if err != nil {
		return result, ErrFailedConversion
	}

	if !ok {
		return result, ErrFailedCasting
	}

	return result, nil
}

func MustParse[T Allowed](value string) T {
	result, err := Parse[T](value)

	if err != nil {
		panic(err)
	}

	return result
}
