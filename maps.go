package gocrazy

import "strings"

/*
GetNestedField returns the value of a nested field in a map.

Usage:

	m := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": map[string]interface{}{
				"baz": "hello",
			},
		},
	}

	val, ok := GetNestedField(m, "foo.bar.baz")
	// val == "hello"
	// ok == true
*/
func GetNestedField(genericMap map[string]interface{}, path string) (interface{}, bool) {
	current := genericMap
	keys := strings.Split(path, ".")
	for i, key := range keys {
		if val, ok := current[key]; ok {
			if i == len(keys)-1 {
				return val, true
			}

			if next, ok := val.(map[string]interface{}); ok {
				current = next
			} else {
				return nil, false
			}
		} else {
			return nil, false
		}
	}
	return nil, false
}
