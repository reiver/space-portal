package cfg

import "github.com/caddyserver/certmagic"

// CertificateAuthority returns the certificate authority endpoint
func CertificateAuthority() string {
	return certmagic.LetsEncryptProductionCA
}
