package helpers

import "fmt"

func ConvertStringSliceToInterfaceSlice(t []string) []interface{} {
	s := make([]interface{}, len(t))

	for i, v := range t {
		s[i] = v
	}

	return s
}

func ConvertFromInterfaceSliceToStringSlice(val interface{}) []string {
	slice := val.([]interface{})
	result := make([]string, len(slice))

	for i, v := range slice {
		result[i] = fmt.Sprint(v)
	}

	return result
}
