package main

import (
	logsrv "github.com/reiver/space-portal/srv/log"
)

func main() {
	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Inform("space-portal ⚡")
	blur()

	const httpstcpaddr string = ":443"
	httpsdaemon := httpsserve(httpstcpaddr)

	const httptcpaddr string = ":80"
	httpdaemon := httpserve(httptcpaddr)

	{
		var err error
		select {
		case err = <-httpsdaemon:
			log.Errorf("https-daemon lost: %s", err)
		case err = <-httpdaemon:
			log.Errorf("http-daemon lost: %s", err)
		}
		panic(err)
	}
}
