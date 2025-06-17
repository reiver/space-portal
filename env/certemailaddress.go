package env

import (
	"os"
)

var CertEMailAddress string = certEMailAddress()

func certEMailAddress() string {
	return os.Getenv("CERT_EMAIL_ADDRESS")
}
