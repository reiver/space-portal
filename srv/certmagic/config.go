package certmagicsrv

import (
	"github.com/caddyserver/certmagic"
	"github.com/reiver/go-erorr"

	"github.com/reiver/space-portal/srv/log"
)

var Config *certmagic.Config

func init() {
	const err = erorr.Error("problem initializing certmagic-config — nil certmagic-config")

	log := logsrv.Prefix("certmagic-config").Begin()
	defer log.End()

	Config = certmagic.NewDefault()
	if nil == Config {
		log.Error(err)
		panic(err)
	}
}
