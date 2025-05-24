package ipv6randomizer

import (
	"crypto/rand"
	"net"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// init registers the module.
func init() {
	caddy.RegisterModule(RandomIPv6Header{})
}

// RandomIPv6Header is a Caddy module that generates a random IPv6 address on each request
// and sets it into the X-Real-IP and X-Forwarded-For headers.
type RandomIPv6Header struct{}

// CaddyModule returns the Caddy module information.
func (RandomIPv6Header) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.random_ipv6_header",
		New: func() caddy.Module { return new(RandomIPv6Header) },
	}
}

// ServeHTTP implements the caddyhttp.MiddlewareHandler interface.
// It generates a new random IPv6 address for each incoming request and assigns it to the headers.
func (rh RandomIPv6Header) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	randomIP, err := generateRandomIPv6()
	if err == nil {
		// Overwrite (or set) the headers with the newly generated randomized IPv6 address.
		r.Header.Set("X-Real-IP", randomIP)
		r.Header.Set("X-Forwarded-For", randomIP)
	}
	// Continue to the next handler in the chain.
	return next.ServeHTTP(w, r)
}

// generateRandomIPv6 creates a random IPv6 address.
// It generates 16 random bytes and converts them to the canonical IPv6 string format.
func generateRandomIPv6() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	ip := net.IP(bytes)
	return ip.String(), nil
}

// Interface guards to ensure the module satisfies the Caddy interfaces.
var (
	_ caddy.Module                = (*RandomIPv6Header)(nil)
	_ caddyhttp.MiddlewareHandler = (*RandomIPv6Header)(nil)
)
