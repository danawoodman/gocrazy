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
		{
			"handle numerical values",
			struct {
				Age    int     `url:"age"`
				Height float64 `url:"height"`
			}{
				Age:    38,
				Height: 6.1,
			},
			"https://example.com?age=38&height=6.1",
		},
		{
			"handle boolean values",
			struct {
				Admin bool `url:"admin"`
			}{
				Admin: true,
			},
			"https://example.com?admin=true",
		},
		{
			"handle omitting fields with `-`",
			struct {
				Ignore string `url:"-"`
			}{
				Ignore: "apple",
			},
			"https://example.com",
		},
		{
			"handle json field tags",
			struct {
				Fruit string `json:"fruit"`
			}{
				Fruit: "apple",
			},
			"https://example.com?fruit=apple",
		},
		{
			"handle ignoring json field tag",
			struct {
				Redirect string `json:"-"`
			}{
				Redirect: "/",
			},
			"https://example.com",
		},
		{
			"fallback to field name",
			struct {
				Query string
			}{
				Query: "some search",
			},
			"https://example.com?Query=some+search",
		},
		{
			"should omit unexported fields",
			struct {
				Exported   string `url:"exported"`
				unexported string `url:"unexported"`
			}{
				Exported:   "a",
				unexported: "b",
			},
			"https://example.com?exported=a",
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
