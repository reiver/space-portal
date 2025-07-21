package main

import (
	"net/http"

	"github.com/reiver/space-portal/lib/middleware"
)

func middlewares(next http.Handler) http.Handler {
	return middleware.Proxy(middleware.ProxyHeaders(next))
}
