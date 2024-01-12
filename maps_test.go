package gocrazy_test

import (
	"testing"

	"github.com/danawoodman/gocrazy"
)

func TestGetNestedField(t *testing.T) {
	type Input struct {
		obj  map[string]interface{}
		path string
	}
	type Output struct {
		val interface{}
		ok  bool
	}
	var tests = []struct {
		name  string
		input Input
		want  Output
	}{
		{
			"returns matched field if present",
			Input{
				obj: map[string]interface{}{
					"foo": map[string]interface{}{
						"bar": map[string]interface{}{
							"baz": "hello",
						},
					},
				},
				path: "foo.bar.baz",
			},
			Output{
				val: "hello",
				ok:  true,
			},
		},
		{
			"returns false for an empty map",
			Input{
				obj:  map[string]interface{}{},
				path: "foo",
			},
			Output{
				val: nil,
				ok:  false,
			},
		},
		{
			"returns false if field does not exist",
			Input{
				obj: map[string]interface{}{
					"foo": map[string]interface{}{
						"bar": map[string]interface{}{
							"baz": "hello",
						},
					},
				},
				path: "does.not.exist",
			},
			Output{
				val: nil,
				ok:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := gocrazy.GetNestedField(tt.input.obj, tt.input.path)
			if val != tt.want.val {
				t.Errorf("Expected %s, got %s", tt.want.val, val)
			}
			if ok != tt.want.ok {
				t.Errorf("Expected %v, got %v", tt.want.ok, ok)
			}
		})
	}

}
