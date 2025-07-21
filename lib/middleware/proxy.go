package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"sync"

	"github.com/reiver/space-portal/cfg"
)

var proxies = new(sync.Map)

// Proxy proxies the request to space-base
func Proxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, _, _ := net.SplitHostPort(r.Host)
		proxy, ok := proxies.Load(host)
		if !ok {
			proxy = httputil.NewSingleHostReverseProxy(cfg.SpaceBaseAddress()) // use .yml file instead of env variable to configure multiple space-base addresses
			proxies.Store(host, proxy)
		}
		proxy.(http.Handler).ServeHTTP(w, r)
		return
	})
}
