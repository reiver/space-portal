package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/caddyserver/certmagic"
	"github.com/reiver/go-erorr"
	"github.com/reiver/space-portal/lib/handler"
	logsrv "github.com/reiver/space-portal/srv/log"
)

func httpserve(tcpaddr string) <-chan error {
	log := logsrv.Prefix("httpserve").Begin()
	defer log.End()

	log.Informf("serving HTTP on TCP address: %q", tcpaddr)

	ch := make(chan error)
	go _httpserve(ch, tcpaddr)
	log.Inform("http-daemon spawed 😈")
	return ch
}

func _httpserve(ch chan error, tcpaddr string) {
	log := logsrv.Prefix("_httpserve").Begin()
	defer log.End()

	cfg := certmagic.NewDefault()

	// Port 80 is used for http-01 challenge, to obtain or renew the certificates
	httpLn, err := net.Listen("tcp", ":80")
	if err != nil {
		err = erorr.Errorf("problem with serving HTTP on TCP address %q: %w", tcpaddr, err)
		log.Errorf("ERROR: %s", err)
		ch <- err
		return
	}

	tlsConfig := cfg.TLSConfig()
	httpsLn, err := tls.Listen("tcp", tcpaddr, tlsConfig)
	if err != nil {
		httpLn.Close()
		httpLn = nil
		err = erorr.Errorf("problem with serving HTTPS on TCP address %q: %w", tcpaddr, err)
		log.Errorf("ERROR: %s", err)
		ch <- err
		return
	}

	// HTTP server for http-01 challenge
	httpServer := &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}
	if len(cfg.Issuers) > 0 {
		if am, ok := cfg.Issuers[0].(*certmagic.ACMEIssuer); ok {
			httpServer.Handler = am.HTTPChallengeHandler(http.HandlerFunc(handler.HTTPRedirectHandler))
		}
	}

	// HTTPS server for TLS termination and proxying
	// TODO: tweak timeouts if needed
	httpsServer := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler: middlewares(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // TODO: remove this, as this will never be called due to `middleware.Proxy`
			w.Write([]byte("Hello, World!"))
		})),
	}

	go func() {
		err = httpServer.Serve(httpLn)
		if err != nil {
			err = erorr.Errorf("problem with serving HTTP on TCP address %q: %w", tcpaddr, err)
			log.Errorf("ERROR: %s", err)
			ch <- err
			return
		}
	}()

	err = httpsServer.Serve(httpsLn)
	if err != nil {
		err = erorr.Errorf("problem with serving HTTPS on TCP address %q: %w", tcpaddr, err)
		log.Errorf("ERROR: %s", err)
		ch <- err
		return
	}
}
