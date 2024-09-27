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
