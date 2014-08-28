package go_utils

type Float64ConversionError struct {
}

func (e Float64ConversionError) Error() string {
	return "Conversion Error: Could not convert num to float64."
}

func ConvNumToFloat64(a interface{}) (float64, error) {
	switch t := a.(type) {
	case int:
		return float64(t), nil
	case uint:
		return float64(t), nil
	case int8:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case uint8:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	}

	return 0, Float64ConversionError{}
}


type IntConversionError struct {
}

func (e IntConversionError) Error() string {
	return "Conversion Error: Could not convert num to int."
}

func ConvNumToInt(a interface{}) (int, error) {
	switch t := a.(type) {
	case int:
		return t, nil
	case uint:
		return int(t), nil
	case int8:
		return int(t), nil
	case int16:
		return int(t), nil
	case int32:
		return int(t), nil
	case int64:
		return int(t), nil
	case uint8:
		return int(t), nil
	case uint16:
		return int(t), nil
	case uint32:
		return int(t), nil
	case uint64:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	}

	return 0, IntConversionError{}
}
