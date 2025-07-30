package middleware

import (
	"net"
	"net/http"

	"github.com/reiver/go-http500"
)

// ProxyHeaders sets the proxy headers before proxying the request
func ProxyHeaders(next http.Handler) http.HandlerFunc {
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

		rip, rport, err := net.SplitHostPort(r.RemoteAddr)
		if nil != err {
			http500.InternalServerError(w, r)
			return
		}

		host, _, err := net.SplitHostPort(r.Host)
		if nil != err {
			http500.InternalServerError(w, r)
			return
		}

		r.Header.Set("Host", host)
		r.Header.Set("X-Forwarded-Proto", "https")
		r.Header.Set("X-Real-IP", rip)
		r.Header.Set("X-Forwarded-For", rip)
		r.Header.Set("X-Forwarded-Port", rport)

		next.ServeHTTP(w, r)
	})
}
