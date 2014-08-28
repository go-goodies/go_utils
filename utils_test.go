package go_utils

import (
	"testing"
	"github.com/remogatto/prettytest"
	"time"
	"errors"
)

type mySuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	prettytest.Run(t, new(mySuite))
}

func (s *mySuite) TestIif() {
	s.Equal(iif_string(true, "true", "false"), "true")
}

func (s *mySuite) TestJoin() {
	data := []string{"A", "B", "C"}
	s.Equal(join(data), "A,B,C")
	s.Equal(join(data, "|"), "A|B|C")
}

func (s *mySuite) TestSubstr() {
	data := "ABCD"
	s.Equal(substr(data, 1, 2), "BC")
}

func (s *mySuite) TestPadRight() {
	data := "ABC"
	s.Equal(padRight(data, 2), "ABC  ")
}

func (s *mySuite) TestPadLeft() {
	data := "ABC"
	s.Equal(padLeft(data, 3), "   ABC")}

type myType struct {
	name string
}

func (s *mySuite) TestTypeOf() {
	s.Equal(typeOf("1"), "string")
	s.Equal(typeOf(int(1)), "int")
	s.Equal(typeOf(1), "int")
	s.Equal(typeOf(uint8(1)), "uint8")
	s.Equal(typeOf(float64(1)), "float64")
	s.Equal(typeOf(1.0), "float64")
	s.Equal(typeOf(time.Now()), "time.Time")

	t, _ := time.Parse("2006-01-02 15:04", "2011-01-19 22:15")
	s.Equal(typeOf(t), "time.Time")

	mt := new(myType)
	mt.name = "alice"
	s.Equal(typeOf(mt), "*go_utils.myType")
}

func (s *mySuite) TestToString() {
	s.Equal(toString(1), "1")
	s.Equal(toString(true), "true")

	// "2014-01-19 22:15:01 -0500" -> 2014-01-19 22:15:01 -0500 EST    ... but
	// "2014-08-28 21:30:01 -0500" -> 2014-08-28 21:30:01 -0500 -0500
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2014-08-28 21:30:01 -0500")
	s.Equal(typeOf(t), "time.Time")
	s.Equal(toString(t), "2014-08-28 21:30:01 -0500 -0500")
}

func (s *mySuite) TestTypes() {
	s.Equal(isBool(true), true)
	s.Equal(isBool(false), true)
	s.Equal(isBool("true"), false)

	s.Equal(isNumber(1), true)
	s.Equal(isNumber(1.0), true)
	s.Equal(isNumber("1.0"), false)
	s.Equal(isNumber(toInt64(1.0)), true)
	s.Equal(isNumber(toInt64("1")), true)
	s.Equal(isNumber(toInt64("1.0")), true)

	s.Equal(isInteger(-1), true)
	s.Equal(isInteger("1"), false)

	s.Equal(isUnsignedInteger(uint(1)), true)
	s.Equal(isUnsignedInteger(-1), false)
	s.Equal(isInteger("1"), false)

	s.Equal(isFloat(1.0), true)

	s.Equal(toInt64("1"), int64(1))

	s.Equal(toUnsignedInteger(1), uint64(1))

	s.Equal(toFloat(1), float64(1.0))

	s.Equal(isError(errors.New("ERROR - List is empty")), true)

	m := make(map[string]string)
	m["1"] = "one"
	s.Equal(isMap(m), true)

	var a [1]string
	a[0] = "Hello"
	s.Equal(isArray(a), true)
	p := []int{2, 3, 5}
	s.Equal(isSlice(p), true)

	s.Equal(isString("1.0"), true)
	s.Equal(isString(padLeft("1.0", 2)), true)

	s.Equal(isString(1.0), false)

	s.Equal(typeOf(1.0), "float64")
	s.Equal(typeOf("1.0"), "string")

	s.Equal(toString(1), "1")

	s.Equal(lengthOf(p), 3)
	s.Equal(lengthOf("ABC"), 3)
	s.Equal(lengthOf(m), 1)

	s.Equal(isNil(nil), true)
	s.Equal(isEmpty(m["XXX"]), true)
}
