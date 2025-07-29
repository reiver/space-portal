package main

import (
	"github.com/reiver/space-portal/cfg"
	logsrv "github.com/reiver/space-portal/srv/log"
	"github.com/reiver/space-portal/srv/tls"
)

func main() {
	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Inform("space-portal ⚡")
	blur()

	// Set default values for the certmagic package (certmagic is used to obtain/renew TLS certificates)
	err := tls.Defaults(cfg.CertEMailAddress(), cfg.CertificateAuthority())
	if err != nil {
		log.Errorf("problem with setting default values for the certmagic package: %s", err)
		panic(err)
	}

	// Don't use port 80, it is reserved for http-01 challenge
	const tcpaddr string = ":443"
	httpdaemon := httpserve(tcpaddr)

	select {
	case err = <-httpdaemon:
		log.Errorf("http-daemon lost: %s", err)
	}
	panic(err)
}
