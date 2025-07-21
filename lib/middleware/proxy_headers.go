package middleware

import (
	"net"
	"net/http"
)

// ProxyHeaders sets the proxy headers before proxying the request
func ProxyHeaders(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		rip, rport, _ := net.SplitHostPort(r.RemoteAddr)
		host, _, _ := net.SplitHostPort(r.Host)

		r.Header.Set("Host", host)
		r.Header.Set("X-Forwarded-Proto", "https")
		r.Header.Set("X-Real-IP", rip)
		r.Header.Set("X-Forwarded-For", rip)
		r.Header.Set("X-Forwarded-Port", rport)

		next.ServeHTTP(w, r)
	})
}
