package main

import (
	"net"
	"net/http"
	"time"

	"github.com/caddyserver/certmagic"
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-http2https"

	"github.com/reiver/space-portal/srv/certmagic"
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

func _httpserve(ch chan<- error, tcpaddr string) {
	log := logsrv.Prefix("_httpserve").Begin()
	defer log.End()

	var listener net.Listener
	{
		var err error

		listener, err = net.Listen("tcp", tcpaddr)
		if err != nil {
			listener.Close()
			listener = nil
			err = erorr.Errorf("problem with serving HTTP on TCP address %q: %w", tcpaddr, err)
			log.Errorf("ERROR: %s", err)
			ch <- err
			return
		}
		if nil == listener {
			err = erorr.Errorf("problem with serving HTTP on TCP address %q — nil listener", tcpaddr)
			log.Errorf("ERROR: %s", err)
			ch <- err
			return
		}
	}

	// HTTP server for http-01 challenge
	var httpServer = http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}
	if len(certmagicsrv.Config.Issuers) > 0 {
		if am, casted := certmagicsrv.Config.Issuers[0].(*certmagic.ACMEIssuer); casted {
			httpServer.Handler = am.HTTPChallengeHandler(http2https.Handler)
		}
	}

	go func() {
		err := httpServer.Serve(listener)
		if err != nil {
			err = erorr.Errorf("problem with serving HTTP on TCP address %q: %w", tcpaddr, err)
			log.Errorf("ERROR: %s", err)
			ch <- err
			return
		}
	}()
}
