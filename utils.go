package go_utils

// Iif_string is and immediate if helper that takes a boolean expression
// and returns a string, true_val if the expression is true else false_val
func Iif_string(expr bool, true_val string, false_val string) string {
	return map[bool]string{true: true_val, false: false_val}[expr]
}
