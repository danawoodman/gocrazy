package gocrazy_test

import (
	"net/url"
	"testing"

	"github.com/danawoodman/gocrazy"
)

func TestAppendQueryParams(t *testing.T) {
	var tests = []struct {
		name  string
		input any
		want  string
	}{
		// the table itself
		{"handle nil", nil, "https://example.com"},
		{
			"handle simple struct",
			struct {
				FirstName string `url:"first_name"`
				LastName  string `url:"last_name"`
			}{
				FirstName: "Gladys",
				LastName:  "Kravitz",
			},
			"https://example.com?first_name=Gladys&last_name=Kravitz",
		},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &url.URL{
				Scheme: "https",
				Host:   "example.com",
			}
			ans := gocrazy.AppendQueryParams(u, tt.input)
			if ans.String() != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
