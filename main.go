package cookiejar

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
)

type CookieJar struct {
	jar     *cookiejar.Jar
	cookies map[*url.URL][]*http.Cookie
	sync.RWMutex
}

func New() *CookieJar {
	jar, _ := cookiejar.New(nil)

	return &CookieJar{
		jar:     jar,
		cookies: map[*url.URL][]*http.Cookie{},
	}
}

func (jar *CookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.Lock()
	defer jar.Unlock()
	jar.cookies[u] = cookies
	jar.jar.SetCookies(u, cookies)
}

func (jar *CookieJar) Cookies(u *url.URL) []*http.Cookie {
	return jar.jar.Cookies(u)
}

func (jar *CookieJar) All() map[*url.URL][]*http.Cookie {
	jar.RLock()
	defer jar.RUnlock()

	cookies := make(map[*url.URL][]*http.Cookie)

	for u, c := range jar.cookies {
		cookies[u] = c
	}

	return cookies
}
