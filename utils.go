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
// Ex:  data: {[A B C]}; Join(data) ==> "A,B,C" ; Join(data, "|") ==> "A|B|C"
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

// padRight pads a string (s) with with a specified string (optional parameter) for padLen characters
// If no string argument is passed, then s will be padded, to the right, with a single space character
func padRight(s string, padLen int, args ...interface{}) string{
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

// padLeft pads a string (s) with with a specified string (optional parameter) for padLen characters
// If no string argument is passed, then s will be padded, to the left, with a single space character
func padLeft(s string, padLen int, args ...interface{}) string{
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

// TypeOf returns the reflection Type of the value in the interface{}.
// TypeOf(nil) returns nil.
func TypeOf(i interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(i))
}

// ToString converts the value to a string
func ToString(value interface{}) string {
	return fmt.Sprint(value)
}
