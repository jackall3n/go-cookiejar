package cookiejar

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	u "net/url"
	"testing"
)

func TestAddCookies(t *testing.T) {
	client := createClient()
	jar := client.Jar.(*CookieJar)

	url := createUrl("https://example.com")

	cookie := &http.Cookie{
		Name:  "some-name",
		Value: "some-value",
	}

	jar.SetCookies(url, []*http.Cookie{cookie})

	assert.Contains(t, jar.Cookies(url), cookie)
}

func TestAllCookies(t *testing.T) {
	client := createClient()
	jar := client.Jar.(*CookieJar)

	url := createUrl("https://example.com")

	cookie := &http.Cookie{
		Name:  "some-name",
		Value: "some-value",
	}

	jar.SetCookies(url, []*http.Cookie{cookie})

	assert.Contains(t, jar.All()[url], cookie)
}

func TestMultipleCookies(t *testing.T) {
	client := createClient()
	jar := client.Jar.(*CookieJar)

	url1 := createUrl("https://example.com")
	url2 := createUrl("https://jck.dev")

	cookie1 := &http.Cookie{
		Name:  "some-name",
		Value: "some-value",
	}

	cookie2 := &http.Cookie{
		Name:  "some-name",
		Value: "some-value",
	}

	cookie3 := &http.Cookie{
		Name:  "some-name",
		Value: "some-value",
	}

	jar.SetCookies(url1, []*http.Cookie{cookie1, cookie2})
	jar.SetCookies(url2, []*http.Cookie{cookie3})

	assert.Contains(t, jar.All()[url1], cookie1)
	assert.Contains(t, jar.All()[url1], cookie2)
	assert.Contains(t, jar.All()[url2], cookie3)
}

// createClient
func createClient() *http.Client {
	client := &http.Client{}
	client.Jar = New()

	return client
}

// createUrl assumes there's no error and returns a URL, removing the tuple
func createUrl(raw string) *u.URL {
	url, _ := u.Parse(raw)

	return url
}
