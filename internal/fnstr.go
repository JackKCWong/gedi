package internal

import (
	"fmt"
	"reflect"
)

func Sprintf(a any, format string) string {
	// check if a is a slice
	ty := reflect.TypeOf(a).Kind()
	if ty == reflect.Slice || ty == reflect.Array {
		// convert a to a slice of strings
		slice := reflect.ValueOf(a)
		args := make([]any, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			args[i] = slice.Index(i).Interface()
		}
		return fmt.Sprintf(format, args...)
	} else {
		return fmt.Sprintf(format, a)
	}
}

func quote(a any, q rune) any {
	// check if a is a slice
	ty := reflect.TypeOf(a).Kind()
	if ty == reflect.Slice || ty == reflect.Array {
		// convert a to a slice of strings
		slice := reflect.ValueOf(a)
		args := make([]any, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			args[i] = fmt.Sprintf("%c%s%c", q, slice.Index(i).Interface(), q)
		}
		return args
	} else {
		return fmt.Sprintf("%c%s%c", q, a, q)
	}
}

func SingleQuote(a any) any {
	return quote(a, '\'')
}

func DoubleQuote(a any) any {
	return quote(a, '"')
}
