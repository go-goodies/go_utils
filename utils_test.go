package go_utils

import (
	"testing"
	"github.com/remogatto/prettytest"
	"time"
)

type mySuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	prettytest.Run(t, new(mySuite))
}

func (s *mySuite) TestIif() {
	s.Equal(Iif_string(true, "true", "false"), "true")
}

func (s *mySuite) TestJoin() {
	data := []string{"A", "B", "C"}
	s.Equal(Join(data), "A,B,C")
	s.Equal(Join(data, "|"), "A|B|C")
}

func (s *mySuite) TestSubstr() {
	data := "ABCD"
	s.Equal(Substr(data, 1, 2), "BC")
}

func (s *mySuite) TestPadRight() {
	data := "ABC"
	s.Equal(PadRight(data, 2), "ABC  ")
}

func (s *mySuite) TestPadLeft() {
	data := "ABC"
	s.Equal(PadLeft(data, 3), "   ABC")}

type myType struct {
	name string
}

func (s *mySuite) TestTypeOf() {
	s.Equal(TypeOf("1"), "string")
	s.Equal(TypeOf(int(1)), "int")
	s.Equal(TypeOf(1), "int")
	s.Equal(TypeOf(uint8(1)), "uint8")
	s.Equal(TypeOf(float64(1)), "float64")
	s.Equal(TypeOf(1.0), "float64")
	s.Equal(TypeOf(time.Now()), "time.Time")

	t, _ := time.Parse("2006-01-02 15:04", "2011-01-19 22:15")
	s.Equal(TypeOf(t), "time.Time")

	mt := new(myType)
	mt.name = "alice"
	s.Equal(TypeOf(mt), "*go_utils.myType")
}

func (s *mySuite) TestToString() {
	s.Equal(ToString(1), "1")
	s.Equal(ToString(true), "true")

	// "2014-01-19 22:15:01 -0500" -> 2014-01-19 22:15:01 -0500 EST    ... but
	// "2014-08-28 21:30:01 -0500" -> 2014-08-28 21:30:01 -0500 -0500
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2014-08-28 21:30:01 -0500")
	s.Equal(TypeOf(t), "time.Time")
	s.Equal(ToString(t), "2014-08-28 21:30:01 -0500 -0500")
}

