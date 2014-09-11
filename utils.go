package go_utils

import (
	"fmt"
	"strings"
	"reflect"
)

// Iif_string is and immediate if helper that takes a boolean expression
// and returns a string, true_val if the expression is true else false_val
func Iif_string(expr bool, true_val string, false_val string) string {
	return map[bool]string{true: true_val, false: false_val}[expr]
}

// Concat joins the strings in a slice, delimiting them with a comma, but it
// allows you to pass the delimiter string to create a single string
// Ex:  data: []string{"A", "B", "C"}; Join(data) ==> "A,B,C" ; Join(data, "|") ==> "A|B|C"
func Join(slice []string, args ...interface{}) string {
	delimiter := ","
	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			delimiter = t
		default:
			panic(fmt.Sprintf("ERROR - Invalid argument (%v).  Must be a string.", arg))
		}
	}
	ret := ""
	for i, s := range slice {
		// append delimiter except at the very end
		ret += s + Iif_string((i < len(slice) - 1), delimiter, "")
	}
	return ret
}

// Substr returns a portion (length characters) of string (s), beginning at a specified position (pos)
func Substr(s string, pos, length int) string{
	runes:=[]rune(s)
	l := pos+length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// PadRight pads a string (s) with with a specified string (optional parameter) for padLen characters
// If no string argument is passed, then s will be padded, to the right, with a single space character
func PadRight(s string, padLen int, args ...interface{}) string{
	padStr := " "
	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			padStr = t
		default:
			panic("Unknown argument")
		}
	}
	return s + strings.Repeat(padStr, padLen);
}

// PadLeft pads a string (s) with with a specified string (optional parameter) for padLen characters
// If no string argument is passed, then s will be padded, to the left, with a single space character
func PadLeft(s string, padLen int, args ...interface{}) string{
	padStr := " "
	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			padStr = t
		default:
			panic("Unknown argument")
		}
	}
	return strings.Repeat(padStr, padLen) + s;
}

// reflect doesn't consider 0 or "" to be zero, so we double check those here
func IsEmpty(args ...interface{}) bool {
	val := reflect.ValueOf(args[0])
	valType := val.Kind()
	switch valType {
	case reflect.String:
		return val.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Interface, reflect.Slice, reflect.Ptr, reflect.Map, reflect.Chan, reflect.Func:
		if val.IsNil() {
			return true
		} else if valType == reflect.Slice || valType == reflect.Map {
			return val.Len() == 0
		}
	case reflect.Struct:
		// verify that all of the struct's properties are empty
		fieldCount := val.NumField()
		for i := 0; i < fieldCount; i++ {
			field := val.Field(i)
			if field.IsValid() && !IsEmpty(field) {
				return false
			}
		}
		return true
	default:
		return false
	}
	return false
}

// Repeat a character (typically used for simple formatting of output)
func Dashes(repeatCount int, args ...interface{}) string{
	dashChar := "-"
	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			dashChar = t
		default:
			panic("Unknown argument")
		}
	}
	return strings.Repeat(dashChar, repeatCount);
}
