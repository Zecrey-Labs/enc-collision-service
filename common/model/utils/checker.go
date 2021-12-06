package utils

import (
	"errors"
	"reflect"
	"strings"
	"unicode"
)

var (
	ErrReflectInvalid     = errors.New("reflect invalid")
	ErrReflectIntInvalid  = errors.New("reflect invalid int type")
	ErrReflectTypeInvalid = errors.New("reflect invalid type")

	ErrEncDataInvalid = errors.New("invalid account index")
)

func CheckRequestParam(dataType uint8, v reflect.Value) error {
	var (
		dataString string
	)
	switch v.Kind() {
	case reflect.Invalid:
		return ErrReflectInvalid
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return ErrReflectIntInvalid
	case reflect.String:
		dataString = v.String()
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return ErrReflectTypeInvalid
	}

	switch dataType {
	case TypeEncData:
		if len(dataString) > maxEncDataLength {
			return ErrEncDataInvalid
		}
		break
	case TypeEncDataOmitSpace:
		if len(dataString) > maxEncDataLengthOmitSpace {
			return ErrEncDataInvalid
		}
		break
	}
	return nil
}

func OmitSpace(s string) string {
	return strings.TrimSpace(s)
}

func OmitSpaceMiddle(s string) (rs string) {
	for _, v := range strings.FieldsFunc(s, unicode.IsSpace) {
		rs = rs + v
	}
	return rs
}

func CleanEncData(encData string) string {
	encData = OmitSpace(encData)
	encData = OmitSpaceMiddle(encData)
	return encData
}
