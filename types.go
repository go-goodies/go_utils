package go_utils

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"strconv"
)
//------------------------------------------------
//               Type Conversions
//------------------------------------------------

// toInt64 converts the argument to a int64
// throws panic if argument is non-numeric parameter
func toInt64(a interface{}) int64 {
	switch t := a.(type) {
	case string:
		if strings.Contains(t, ".") {
			fpNum, err := strconv.ParseFloat(t, 64)
			if err != nil {
				log.Println(StringConversionError{})
				return 0
			}
			return int64(fpNum)
		} else {
			num, err := ConvStringToInt(a)
			if err != nil {
				panic(fmt.Sprintf("Error converting string!  Got <%T> %#v", a, a))
			}
			return int64(num)
		}
	}
	if isInteger(a) {
		return reflect.ValueOf(a).Int()
	} else if isUnsignedInteger(a) {
		return int64(reflect.ValueOf(a).Uint())
	} else if isFloat(a) {
		return int64(reflect.ValueOf(a).Float())
	} else {
		panic(fmt.Sprintf("Expected a number!  Got <%T> %#v", a, a))
	}
}

// toFloat converts the argument to a float64
// throws panic if argument is non-numeric parameter
func toFloat(a interface{}) float64 {
	if isInteger(a) {
		return float64(reflect.ValueOf(a).Int())
	} else if isUnsignedInteger(a) {
		return float64(reflect.ValueOf(a).Uint())
	} else if isFloat(a) {
		return reflect.ValueOf(a).Float()
	} else {
		panic(fmt.Sprintf("Expected a number!  Got <%T> %#v", a, a))
	}
}

// toUnsignedInteger converts the argument to a uint64
// throws panic if argument is non-numeric parameter
func toUnsignedInteger(a interface{}) uint64 {
	if isInteger(a) {
		return uint64(reflect.ValueOf(a).Int())
	} else if isUnsignedInteger(a) {
		return reflect.ValueOf(a).Uint()
	} else if isFloat(a) {
		return uint64(reflect.ValueOf(a).Float())
	} else {
		panic(fmt.Sprintf("Expected a number!  Got <%T> %#v", a, a))
	}
}

//------------------------------------------------
//               Type Equality
//------------------------------------------------

func isArray(a interface{}) bool {
	if a == nil {
		return false
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Array:
		return true
	default:
		return false
	}
}

func isBool(a interface{}) bool {
	return reflect.TypeOf(a).Kind() == reflect.Bool
}

func isChan(a interface{}) bool {
	if isNil(a) {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.Chan
}

func isError(a interface{}) bool {
	_, ok := a.(error)
	return ok
}

func isFloat(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return kind >= reflect.Float32 && kind <= reflect.Float64
}

func isInteger(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return kind >= reflect.Int && kind <= reflect.Int64
}

func isMap(a interface{}) bool {
	if a == nil {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.Map
}

func isNil(a interface{}) bool {
	if a == nil {
		return true
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(a).IsNil()
	}
	return false
}

func isNumber(a interface{}) bool {
	if a == nil {
		return false
	}
	kind := reflect.TypeOf(a).Kind()
	return kind >= reflect.Int  && kind <= reflect.Float64
}

func isSlice(a interface{}) bool {
	if a == nil {
		return false
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Slice:
		return true
	default:
		return false
	}
}

func isString(a interface{}) bool {
	if a == nil {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.String
}

func isUnsignedInteger(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return kind >= reflect.Uint && kind <= reflect.Uint64
}

//------------------------------------------------
//                 Type Helpers
//------------------------------------------------

// lengthOf returns the number if items in argument
func lengthOf(a interface{}) int {
	if a == nil {
		return 0
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Chan, reflect.Array:
		return reflect.ValueOf(a).Len()
	default:
		return 0
	}
}

// toString converts the value to a string
func toString(a interface{}) string {
	return fmt.Sprint(a)
}

// typeOf returns the reflection Type of the value in the interface{}.
// typeOf(nil) returns nil.
func typeOf(a interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(a))
}
