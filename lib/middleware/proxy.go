package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/reiver/space-portal/cfg"
)

var proxies = make(map[string]*httputil.ReverseProxy)

// Proxy proxies the request to space-base
func Proxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, _, _ := net.SplitHostPort(r.Host)
		proxy, ok := proxies[host]
		if !ok {
			proxy = httputil.NewSingleHostReverseProxy(cfg.SpaceBaseAddress()) // use .yml file instead of env variable to configure multiple space-base addresses
			proxies[host] = proxy
		}
		proxy.ServeHTTP(w, r)
		return
	})
}
