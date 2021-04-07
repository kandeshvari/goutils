package goutils

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

var urlCases = []struct {
	URL string
	Res *URL
}{
	{
		URL: "https://www.google.com/search?newwindow=1&client=ubuntu&hs=2w9&channel=fs&q=rubenv+sql-migrate+%22clickhouse%22&sa=X&ved=2ahUKEwiuhP-eg-zvAhVQEncKHfe7DXMQ5t4CMAB6BAgEEA4&biw=1470&bih=744",
		Res: &URL{
			Schema: "https",
			Host:   "www.google.com",
			Path:   "/search",
			Query: url.Values{
				"newwindow": []string{"1"},
				"client":    []string{"ubuntu"},
				"hs":        []string{"2w9"},
				"channel":   []string{"fs"},
				"q":         []string{`rubenv sql-migrate "clickhouse"`},
				"sa":        []string{"X"},
				"ved":       []string{"2ahUKEwiuhP-eg-zvAhVQEncKHfe7DXMQ5t4CMAB6BAgEEA4"},
				"biw":       []string{"1470"},
				"bih":       []string{"744"},
			},
		},
	},
	{
		URL: "http://www.google.com/",
		Res: &URL{
			Schema: "http",
			Host:   "www.google.com",
			Path:   "/",
			Query:  url.Values{},
		},
	},
	{
		URL: "http://www.google.com/p/q/r/s?q=1#anchor",
		Res: &URL{
			Schema: "http",
			Host:   "www.google.com",
			Path:   "/p/q/r/s",
			Hash:   "anchor",
			Query: url.Values{
				"q": []string{"1"},
			},
		},
	},
}

func TestParseURL(t *testing.T) {
	for _, c := range urlCases {
		assert.Equal(t, c.Res, ParseURL(c.URL), "url parsed incorrectly")
	}
}
