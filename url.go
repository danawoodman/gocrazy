package gocrazy

import (
	"fmt"
	"net/url"
	"reflect"
)

func AppendQueryParams(u *url.URL, paramData any) *url.URL {
	q := u.Query()
	val := reflect.ValueOf(paramData)

	// If the value is a pointer, we need to get the value of the pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		name := field.Tag.Get("url")
		if name != "" {
			q.Set(name, fmt.Sprintf("%v", val.Field(i).Interface()))
		}
	}
	u.RawQuery = q.Encode()

	return u
}
