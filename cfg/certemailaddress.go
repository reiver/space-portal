package cfg

import (
	"github.com/reiver/space-portal/env"
)

// space-portal deals with all things related to the HTTP TLS certificates for the space-bases, so that the space-bases don't have to.
//
// When an HTTP TLS certificate is created, it needs to have an e-mail address associated with it.
//
// This e-mail address for sending expiration and other important notifications.
// This e-mail address is not embedded within the certificate's data.
func CertEMailAddress() string {
	return env.CertEMailAddress
}
