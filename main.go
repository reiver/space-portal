package main

import (
	"github.com/reiver/space-portal/srv/log"
)

func main() {
	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Inform("space-portal ⚡")
	blur()

	const tcpaddr string = ":80"
	httpdaemon := httpserve(tcpaddr)

	var err error
	select {
	case err = <-httpdaemon:
		log.Errorf("http-daemon lost: %s", err)
	}
	panic(err)
}
