package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/reiver/go-erorr"

	"github.com/reiver/space-portal/srv/certmagic"
	"github.com/reiver/space-portal/srv/log"
)

func httpsserve(tcpaddr string) <-chan error {
	log := logsrv.Prefix("httpsserve").Begin()
	defer log.End()

	log.Informf("serving HTTPS on TCP address: %q", tcpaddr)

	ch := make(chan error)
	go _httpserve(ch, tcpaddr)
	log.Inform("https-daemon spawed 😈")
	return ch
}

func _httpsserve(ch chan<- error, tcpaddr string) {
	log := logsrv.Prefix("_httpsserve").Begin()
	defer log.End()

	var listener net.Listener
	{
		var err error

		tlsConfig := certmagicsrv.Config.TLSConfig()
		listener, err = tls.Listen("tcp", tcpaddr, tlsConfig)
		if err != nil {
			listener.Close()
			listener = nil
			err = erorr.Errorf("problem with serving HTTPS on TCP address %q: %w", tcpaddr, err)
			log.Errorf("ERROR: %s", err)
			ch <- err
			return
		}
		if nil == listener {
			err = erorr.Errorf("problem with serving HTTPS on TCP address %q — nil listener", tcpaddr)
			log.Errorf("ERROR: %s", err)
			ch <- err
			return
		}
	}

	var handler http.Handler
	{
		handler = http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) { // TODO: remove this, as this will never be called due to `middleware.Proxy`
				w.Write([]byte("Hello, World!"))
			},
		)

		handler = middlewares(handler)

		if nil == handler {
			err := erorr.Errorf("problem with serving HTTPS on TCP address %q — could not set-up http-handler and http-middleware", tcpaddr)
			log.Errorf("ERROR: %s", err)
			ch <- err
			return
		}
	}

	// HTTPS server for TLS termination and proxying
	// TODO: tweak timeouts if needed
	var httpsServer = http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler: handler,
	}

	go func() {
		err := httpsServer.Serve(listener)
		if err != nil {
			err = erorr.Errorf("problem with serving HTTPS on TCP address %q: %w", tcpaddr, err)
			log.Errorf("ERROR: %s", err)
			ch <- err
			return
		}
	}()
}
