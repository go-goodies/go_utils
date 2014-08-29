package go_utils

// Check whether the expected (e) string is in the slice (s) of stings
func StringSliceContains(e string, s []string) bool {
	for _, a := range s { if a == e { return true } }
	return false
}

// Check whether the expected (e) int is in the slice (s) of ints
func IntSliceContains(e int, s []int) bool {
	for _, a := range s { if a == e { return true } }
	return false
}

// Check whether the expected (e) string is in the slice (s) of stings
func Float64SliceContains(e float64, s []float64) bool {
	for _, a := range s { if a == e { return true } }
	return false
}

// ----------------------
//     Custom Types
// ----------------------

type Widget struct {
	Name	string	`json:"name,omitempty"`		// the field is omitted from the object if its value is empty
	Count	int64	`json:"count,omitempty"` 	// the field is omitted from the object if its value is empty
}

// Check whether the expected (e) widget is in the slice (s) of widgets
func ContainsWidget(e *Widget, widgets []*Widget) bool {
	for _, d := range widgets {
		if d.Name == e.Name &&
			d.Count == e.Count {
			return true
		}
	}
	return false
}
