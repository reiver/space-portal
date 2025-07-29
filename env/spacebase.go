package env

import (
	"net/url"
	"os"

	logsrv "github.com/reiver/space-portal/srv/log"
)

// SpaceBaseAddress returns the space-base address
var SpaceBaseAddress *url.URL = spaceBaseAddress()

// spaceBaseAddress returns the space-base address from the environment variable
// it should be in the format of "http://{ip|hostname}:{port}"
func spaceBaseAddress() *url.URL {
	spaceBaseAddress := os.Getenv("SPACE_BASE_ADDRESS")
	url, err := url.Parse(spaceBaseAddress)
	if err != nil {
		logsrv.Prefix("env.spaceBaseAddress").Errorf("invalid SPACE_BASE_ADDRESS: %s", err)
		panic(err)
	}
	return url
}
