package tls

import "github.com/caddyserver/certmagic"

// Defaults sets the default values for the certmagic package (certmagic is used to obtain/renew TLS certificates)
func Defaults(certEmail, certCA string) error {
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
