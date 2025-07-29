package cfg

import (
	"net/url"

	"github.com/reiver/space-portal/env"
)

// SpaceBaseAddress returns the space-base address
func SpaceBaseAddress() *url.URL {
	return env.SpaceBaseAddress
}
