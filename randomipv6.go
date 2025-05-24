package randomipv6

import (
	"crypto/rand"
	"net"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	// Register the module.
	caddy.RegisterModule(RandomIPv6{})
	// Register the Caddyfile directive and its parsing function.
	httpcaddyfile.RegisterHandlerDirective("randomipv6", parseCaddyfile)
	httpcaddyfile.RegisterDirectiveOrder("randomipv6", httpcaddyfile.After, "templates")
}

// RandomIPv6 is a Caddy HTTP middleware module that injects a randomized IPv6
// address into the X-Real-IP and X-Forwarded-For headers per request.
type RandomIPv6 struct{}

// CaddyModule returns the Caddy module information.
func (RandomIPv6) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.randomipv6",
		New: func() caddy.Module { return new(RandomIPv6) },
	}
}

// ServeHTTP generates a new random IPv6 address for each request, sets it into
// the required headers, and then passes control to the next handler.
func (r *RandomIPv6) ServeHTTP(w http.ResponseWriter, req *http.Request, next caddyhttp.Handler) error {
	ip, err := generateRandomIPv6()
	if err != nil {
		return err
	}

	req.Header.Set("X-Real-IP", ip)
	req.Header.Set("X-Forwarded-For", ip)

	return next.ServeHTTP(w, req)
}

// generateRandomIPv6 creates a random IPv6 address by reading 16 random bytes.
func generateRandomIPv6() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	ip := net.IP(b)
	return ip.String(), nil
}

// parseCaddyfile is the function responsible for parsing the `randomipv6` directive
// inside a Caddyfile and applying it correctly.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	module := new(RandomIPv6)
	return module, nil
}

// Interface guards
var (
	_ caddy.Module                = (*RandomIPv6)(nil)
	_ caddyhttp.MiddlewareHandler = (*RandomIPv6)(nil)
)
