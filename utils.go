// version 1.0.7
package envstruct

import (
	"errors"
	"reflect"
	"strconv"
)

type parserFunc func(v string) (interface{}, error)

// type defaultFunc func() interface{}
type TypeDefaultBy struct {
	ParserFunc   parserFunc
	ValueDefault interface{}
}

// type map[reflect.Type]

var (
	// ErrNotAStructPtr is returned if you pass something that is not a pointer to a
	// Struct to Parse.
	ErrNotAStructPtr = errors.New("env: expected a pointer to a Struct")

	DefaultByType = map[reflect.Type]TypeDefaultBy{
		reflect.TypeOf(false): {
			ParserFunc: func(v string) (interface{}, error) {
				return strconv.ParseBool(v)
			},
			ValueDefault: bool(false),
		},
		reflect.TypeOf(""): {
			ParserFunc: func(v string) (interface{}, error) {
				return v, nil
			},
			ValueDefault: string(""),
		},
		reflect.TypeOf(int(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				i, err := strconv.ParseInt(v, 10, 32)
				return int(i), err
			},
			ValueDefault: int(0),
		},
		reflect.TypeOf(int8(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				i, err := strconv.ParseInt(v, 10, 8)
				return int8(i), err
			},
			ValueDefault: int8(0),
		},
		reflect.TypeOf(int16(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				i, err := strconv.ParseInt(v, 10, 16)
				return int16(i), err
			},
			ValueDefault: int16(0),
		},
		reflect.TypeOf(int32(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				i, err := strconv.ParseInt(v, 10, 32)
				return int32(i), err
			},
			ValueDefault: int32(0),
		},
		reflect.TypeOf(int64(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				return strconv.ParseInt(v, 10, 64)
			},
			ValueDefault: int64(0),
		},
		reflect.TypeOf(uint(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				i, err := strconv.ParseUint(v, 10, 32)
				return uint(i), err
			},
			ValueDefault: uint(0),
		},
		reflect.TypeOf(uint8(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				i, err := strconv.ParseUint(v, 10, 8)
				return uint8(i), err
			},
			ValueDefault: uint8(0),
		},
		reflect.TypeOf(uint16(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				i, err := strconv.ParseUint(v, 10, 16)
				return uint16(i), err
			},
			ValueDefault: uint(16),
		},
		reflect.TypeOf(uint32(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				i, err := strconv.ParseUint(v, 10, 32)
				return uint32(i), err
			},
			ValueDefault: uint32(0),
		},
		reflect.TypeOf(uint64(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				i, err := strconv.ParseUint(v, 10, 64)
				return i, err
			},
			ValueDefault: uint64(0),
		},
		reflect.TypeOf(float32(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				f, err := strconv.ParseFloat(v, 32)
				return float32(f), err
			},
			ValueDefault: float32(0),
		},
		reflect.TypeOf(float64(0)): {
			ParserFunc: func(v string) (interface{}, error) {
				return strconv.ParseFloat(v, 64)
			},
			ValueDefault: float64(0),
		},
	}
)
