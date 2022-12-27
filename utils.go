package envstruct

import (
	"errors"
	"reflect"
	"strconv"
)

type parserFunc func(v string) (interface{}, error)
type defaultFunc func() interface{}

var (
	// ErrNotAStructPtr is returned if you pass something that is not a pointer to a
	// Struct to Parse.
	ErrNotAStructPtr = errors.New("env: expected a pointer to a Struct")

	defaultBuiltInParsers = map[reflect.Kind]parserFunc{
		reflect.Bool: func(v string) (interface{}, error) {
			return strconv.ParseBool(v)
		},
		reflect.String: func(v string) (interface{}, error) {
			return v, nil
		},
		reflect.Int: func(v string) (interface{}, error) {
			i, err := strconv.ParseInt(v, 10, 32)
			return int(i), err
		},
		reflect.Int8: func(v string) (interface{}, error) {
			i, err := strconv.ParseInt(v, 10, 8)
			return int8(i), err
		},
		reflect.Int16: func(v string) (interface{}, error) {
			i, err := strconv.ParseInt(v, 10, 16)
			return int16(i), err
		},
		reflect.Int32: func(v string) (interface{}, error) {
			i, err := strconv.ParseInt(v, 10, 32)
			return int32(i), err
		},
		reflect.Int64: func(v string) (interface{}, error) {
			return strconv.ParseInt(v, 10, 64)
		},
		reflect.Uint: func(v string) (interface{}, error) {
			i, err := strconv.ParseUint(v, 10, 32)
			return uint(i), err
		},
		reflect.Uint8: func(v string) (interface{}, error) {
			i, err := strconv.ParseUint(v, 10, 8)
			return uint8(i), err
		},
		reflect.Uint16: func(v string) (interface{}, error) {
			i, err := strconv.ParseUint(v, 10, 16)
			return uint16(i), err
		},
		reflect.Uint32: func(v string) (interface{}, error) {
			i, err := strconv.ParseUint(v, 10, 32)
			return uint32(i), err
		},
		reflect.Uint64: func(v string) (interface{}, error) {
			i, err := strconv.ParseUint(v, 10, 64)
			return i, err
		},
		reflect.Float32: func(v string) (interface{}, error) {
			f, err := strconv.ParseFloat(v, 32)
			return float32(f), err
		},
		reflect.Float64: func(v string) (interface{}, error) {
			return strconv.ParseFloat(v, 64)
		},
	}

	defaultValueMap = map[reflect.Kind]defaultFunc{
		reflect.Bool: func() interface{} {
			return false
		},
		reflect.String: func() interface{} {
			return ""
		},
		reflect.Int: func() interface{} {
			return int(0)
		},
		reflect.Int8: func() interface{} {
			return int8(0)
		},
		reflect.Int16: func() interface{} {
			return int16(0)
		},
		reflect.Int32: func() interface{} {
			return int32(0)
		},
		reflect.Int64: func() interface{} {
			return int64(0)
		},
		reflect.Uint: func() interface{} {
			return uint(0)
		},
		reflect.Uint8: func() interface{} {
			return uint8(0)
		},
		reflect.Uint16: func() interface{} {
			return uint16(0)
		},
		reflect.Uint32: func() interface{} {
			return uint32(0)
		},
		reflect.Uint64: func() interface{} {
			return uint64(0)
		},
		reflect.Float32: func() interface{} {
			return float32(0)
		},
		reflect.Float64: func() interface{} {
			return float64(0)
		},
	}
)
