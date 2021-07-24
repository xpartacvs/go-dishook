package dishook

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UrlFormat_CaseInsensitive_Allowed(t *testing.T) {
	urls := []Url{
		"http://domain",
		"https://domain",
		"Http://doMain",
		"Https://doMain",
	}

	for _, url := range urls {
		require.NoError(
			t,
			url.validate(),
			"Insensitive case in %s should be allowed. Please fix func (Url).validate()",
			url,
		)
	}
}

func Test_UrlFormat_Valid(t *testing.T) {
	urls := []Url{
		"http://domain",
		"https://domain",
		"Http://domain",
		"Https://domain",
		"Http://doMain",
		"Https://doMain",
	}

	for _, url := range urls {
		require.NoError(
			t,
			url.validate(),
			"%s should be valid. Please fix func (Url).validate()",
			url,
		)
	}
}
func Test_UrlFormat_Invalid(t *testing.T) {
	urls := []Url{
		"",
		"not-start-with-http-or-https",
		"Http:/ domain",
		"Htts://domain",
		"Htt://doMain",
		"Https://",
		"http://d",
	}

	for _, url := range urls {
		require.Error(
			t,
			url.validate(),
			"%s should be valid. Please fix func (Url).validate()",
			url,
		)
	}
}
