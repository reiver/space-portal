package tlssrv

import (
	"github.com/caddyserver/certmagic"
	"github.com/reiver/go-erorr"

	"github.com/reiver/space-portal/cfg"
	"github.com/reiver/space-portal/srv/log"
)

// defaults sets the default values for the certmagic package (certmagic is used to obtain/renew TLS certificates)
func defaults(certEmail, certCA string) error {
	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = certEmail
	certmagic.DefaultACME.CA = certCA

	// set default on-demand logic
	err := OnDemand()
	if err != nil {
		return err
	}

	return nil
}

func init () {
	log := logsrv.Prefix("tlssrv-defaults-init").Begin()
	defer log.End()

        // Set default values for the certmagic package (certmagic is used to obtain/renew TLS certificates)
	err := defaults(cfg.CertEMailAddress(), cfg.CertificateAuthority())
	if err != nil {
		err = erorr.Errorf("problem with setting default values for the certmagic package: %s", err)
		log.Error(err)
		panic(err)
	}
}
