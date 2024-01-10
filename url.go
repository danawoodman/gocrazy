package gocrazy

import (
	"fmt"
	"net/url"
	"reflect"
)

/*
AppendQueryParams appends the values from the struct to the URL.

The params are determined by the following rules:

- If the field has a `url` tag set, the value of the tag is used as the query parameter name.
- If no `url` tag set, fallback to using a `json` tag if present
- If not `url` or `json` tags are set, fallback to using the field name

Fields are ignored if:

- If it is not exported, it is ignored.
- If it has a `url` (or falling back to `json`) tag set to `-`, it is ignored.
*/
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

		// Don't add unexported fields:
		if field.IsExported() {
			urlTag := field.Tag.Get("url")
			jsonTag := field.Tag.Get("json")
			if urlTag != "" && urlTag != "-" {
				q.Set(urlTag, fmt.Sprintf("%v", fieldValue.Interface()))
				// Fallback to the `json` tag if no `url` tag is set:
			} else if jsonTag != "" && jsonTag != "-" {
				q.Set(jsonTag, fmt.Sprintf("%v", fieldValue.Interface()))
				// Don't use tag name if the user has requested the field be omitted:
			} else if field.Name != "" && urlTag != "-" && jsonTag != "-" {
				q.Set(field.Name, fmt.Sprintf("%v", fieldValue.Interface()))
			}
		}
	}
	u.RawQuery = q.Encode()

	return u
}
