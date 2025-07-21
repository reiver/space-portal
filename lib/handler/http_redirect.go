package handler

import (
	"net"
	"net/http"
)

// Handler for redirecting HTTP traffic to HTTPS after obtaining the certificate
func HTTPRedirectHandler(w http.ResponseWriter, r *http.Request) {
	toURL := "https://"

	host, _, _ := net.SplitHostPort(r.Host)

	toURL += host
	toURL += r.URL.RequestURI()

	w.Header().Set("Connection", "close")

	http.Redirect(w, r, toURL, http.StatusMovedPermanently)
}
