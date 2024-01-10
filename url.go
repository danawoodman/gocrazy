package gocrazy

import (
	"fmt"
	"net/url"
	"reflect"
)

// AppendQueryParams appends the values from the struct to the URL.
// The struct must have a `url` tag on each field that should be appended to the URL.
func AppendQueryParams(u *url.URL, paramData any) *url.URL {
	q := u.Query()
	val := reflect.ValueOf(paramData)

	// If the value is a pointer, we need to get the value of the pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.IsValid() == false {
		return u
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		// Skip zero values
		if fieldValue.IsZero() {
			continue
		}

		name := field.Tag.Get("url")
		if name != "" {
			q.Set(name, fmt.Sprintf("%v", fieldValue.Interface()))
		}
	}
	u.RawQuery = q.Encode()

	return u
}
