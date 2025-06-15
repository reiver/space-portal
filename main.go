package main

import (
	"github.com/reiver/space-portal/srv/log"
)

func main() {
	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Inform("space-portal ⚡")
	blur()
}
