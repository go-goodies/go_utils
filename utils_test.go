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

func (s *mySuite) TestTypes() {
	s.Equal(IsBool(true), true)
	s.Equal(IsBool(false), true)
	s.Equal(IsBool("true"), false)

	s.Equal(IsNumber(1), true)
	s.Equal(IsNumber(1.0), true)
	s.Equal(IsNumber("1.0"), false)
	s.Equal(IsNumber(ToInt64(1.0)), true)
	s.Equal(IsNumber(ToInt64("1")), true)
	s.Equal(IsNumber(ToInt64("1.0")), true)

	s.Equal(IsInteger(-1), true)
	s.Equal(IsInteger("1"), false)

	s.Equal(IsUnsignedInteger(uint(1)), true)
	s.Equal(IsUnsignedInteger(-1), false)
	s.Equal(IsInteger("1"), false)

	s.Equal(IsFloat(1.0), true)

	s.Equal(ToInt64("1"), int64(1))

	s.Equal(ToUnsignedInteger(1), uint64(1))

	s.Equal(ToFloat(1), float64(1.0))

	s.Equal(IsError(errors.New("ERROR - List is empty")), true)

	m := make(map[string]string)
	m["1"] = "one"
	s.Equal(IsMap(m), true)

	var a [1]string
	a[0] = "Hello"
	s.Equal(IsArray(a), true)
	p := []int{2, 3, 5}
	s.Equal(IsSlice(p), true)

	s.Equal(IsString("1.0"), true)
	s.Equal(IsString(PadLeft("1.0", 2)), true)

	s.Equal(IsString(1.0), false)

	s.Equal(TypeOf(1.0), "float64")
	s.Equal(TypeOf("1.0"), "string")

	s.Equal(ToString(1), "1")

	s.Equal(LengthOf(p), 3)
	s.Equal(LengthOf("ABC"), 3)
	s.Equal(LengthOf(m), 1)

	s.Equal(IsNil(nil), true)
	s.Equal(IsEmpty(m["XXX"]), true)
}

func (s *mySuite) TestContains() {

	s.Equal(StringSliceContains("1", []string{"2", "3", "1"}), true)
	s.Equal(StringSliceContains("1", []string{"2", "3", "5"}), false)

	s.Equal(IntSliceContains(1, []int{2, 3, 1}), true)
	s.Equal(IntSliceContains(1, []int{2, 3, 5}), false)

	s.Equal(Float64SliceContains(1.0, []float64{2.0, 3.0, 1.0}), true)
	s.Equal(Float64SliceContains(1.0, []float64{2.0, 3.0, 5.0}), false)

	widgets := []*Widget{}
	widgets = append(widgets, &Widget{"A", 1})
	widgets = append(widgets, &Widget{"B", 2})
	widgets = append(widgets, &Widget{"C", 3})

	widget := new(Widget)
	widget.Name = "B"
	widget.Count = 2
	s.Equal(ContainsWidget(widget, widgets), true)
	widget.Count = 9
	s.Equal(ContainsWidget(widget, widgets), false)
}


func (s *mySuite) TestSingleton() {
	var AppContext *Singleton
	AppContext = NewSingleton()
	AppContext.Data["username"] = "joesample"
	s.Equal(AppContext.Data["username"], "joesample")
}

func (s *mySuite) TestDashes() {

	s.Equal(Dashes(3, "-"), "---")
}

