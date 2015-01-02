package go_utils

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

//------------------------------------------------
//               Type Conversions
//------------------------------------------------

// ToInt64 converts the argument to a int64
// throws panic if argument is non-numeric parameter
func ToInt64(a interface{}) int64 {
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
	if IsInteger(a) {
		return reflect.ValueOf(a).Int()
	} else if IsUnsignedInteger(a) {
		return int64(reflect.ValueOf(a).Uint())
	} else if IsFloat(a) {
		return int64(reflect.ValueOf(a).Float())
	} else {
		panic(fmt.Sprintf("Expected a number!  Got <%T> %#v", a, a))
	}
}

// ToFloat converts the argument to a float64
// throws panic if argument is non-numeric parameter
func ToFloat(a interface{}) float64 {
	if IsInteger(a) {
		return float64(reflect.ValueOf(a).Int())
	} else if IsUnsignedInteger(a) {
		return float64(reflect.ValueOf(a).Uint())
	} else if IsFloat(a) {
		return reflect.ValueOf(a).Float()
	} else {
		panic(fmt.Sprintf("Expected a number!  Got <%T> %#v", a, a))
	}
}

// ToUnsignedInteger converts the argument to a uint64
// throws panic if argument is non-numeric parameter
func ToUnsignedInteger(a interface{}) uint64 {
	if IsInteger(a) {
		return uint64(reflect.ValueOf(a).Int())
	} else if IsUnsignedInteger(a) {
		return reflect.ValueOf(a).Uint()
	} else if IsFloat(a) {
		return uint64(reflect.ValueOf(a).Float())
	} else {
		panic(fmt.Sprintf("Expected a number!  Got <%T> %#v", a, a))
	}
}

//------------------------------------------------
//               Type Equality
//------------------------------------------------

func IsArray(a interface{}) bool {
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

func IsBool(a interface{}) bool {
	return reflect.TypeOf(a).Kind() == reflect.Bool
}

func IsChan(a interface{}) bool {
	if IsNil(a) {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.Chan
}

func IsError(a interface{}) bool {
	_, ok := a.(error)
	return ok
}

func IsFloat(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return kind >= reflect.Float32 && kind <= reflect.Float64
}

func IsInteger(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return kind >= reflect.Int && kind <= reflect.Int64
}

func IsInt(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return kind == reflect.Int
}

func IsUint(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return kind == reflect.Uint
}

func IsMap(a interface{}) bool {
	if a == nil {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.Map
}

func IsNil(a interface{}) bool {
	if a == nil {
		return true
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(a).IsNil()
	}
	return false
}

func IsNumber(a interface{}) bool {
	if a == nil {
		return false
	}
	kind := reflect.TypeOf(a).Kind()
	return kind >= reflect.Int && kind <= reflect.Float64
}

func IsSlice(a interface{}) bool {
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

func IsString(a interface{}) bool {
	if a == nil {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.String
}

func IsUnsignedInteger(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return kind >= reflect.Uint && kind <= reflect.Uint64
}

//------------------------------------------------
//                 Type Helpers
//------------------------------------------------

// LengthOf returns the number if items in argument
func LengthOf(a interface{}) int {
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

// ToString converts the value to a string
func ToString(a interface{}) string {
	ret := ""
	if strings.Index(TypeOf(a), "map[string]string") > -1 {
		// will return nicely formatted results for this type (map[string]string) of map only
		m := a.(map[string]string)
		for key, value := range m {
			ret = fmt.Sprintf("%s: %v, ", strings.TrimSpace(key), strings.TrimSpace(value))
		}
		ret = strings.TrimSuffix(ret, ", ")

	} else {
		ret = fmt.Sprint(a)
	}
	return ret
}

// TypeOf returns the reflection Type of the value in the interface{}.
// TypeOf(nil) returns nil.
func TypeOf(a interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(a))
}
