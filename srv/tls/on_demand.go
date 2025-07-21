package tls

import (
	"context"

	"github.com/caddyserver/certmagic"
)

// OnDemand  will be called when a new TLS certificate is needed
func OnDemand() error {
	certmagic.Default.OnDemand = &certmagic.OnDemandConfig{
		DecisionFunc: func(ctx context.Context, domain string) error {
			// TODO: add logic to decide if a certificate should be generated for the given domain name
			// `return nil` means that the certificate will be generated for the given domain name
			// `return err` means that the certificate will not be generated for the given domain name, and the request will be denied
			return nil
		},
	}
	return nil
}
