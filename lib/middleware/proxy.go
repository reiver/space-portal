package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/reiver/go-http500"
	"github.com/reiver/go-reg"

	"github.com/reiver/space-portal/cfg"
)

var proxies reg.Registry[http.Handler]

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

		proxy, found := proxies.Get(host)
		if !found {
			proxy = httputil.NewSingleHostReverseProxy(cfg.SpaceBaseAddress()) // use .yml file instead of env variable to configure multiple space-base addresses
			proxies.Set(host, proxy)
		}
		proxy.ServeHTTP(w, r)
		return
	})
}
