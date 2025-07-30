package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"sync"

	"github.com/reiver/go-http500"

	"github.com/reiver/space-portal/cfg"
)

var proxies = new(sync.Map)

// Proxy proxies the request to space-base
func Proxy(next http.Handler) http.Handler {
	if nil == next {
		return nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if nil == w {
			return
		}
		if nil == r {
			http500.InternalServerError(w, r)
			return
		}

		host, _, err := net.SplitHostPort(r.Host)
		if nil != err {
			http500.InternalServerError(w, r)
			return
		}

		proxy, found := proxies.Load(host)
		if !found {
			proxy = httputil.NewSingleHostReverseProxy(cfg.SpaceBaseAddress()) // use .yml file instead of env variable to configure multiple space-base addresses
			proxies.Store(host, proxy)
		}
		proxy.(http.Handler).ServeHTTP(w, r)
		return
	})
}
