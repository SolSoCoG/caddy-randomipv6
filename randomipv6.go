package randomipv6

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// IPv6Middleware struct
type IPv6Middleware struct{}

// Caddy module registration
func init() {
	caddy.RegisterModule(IPv6Middleware{})
}

// Returns the module ID
func (IPv6Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.random_ipv6",
		New: func() caddy.Module { return new(IPv6Middleware) },
	}
}

// ServeHTTP function injects a random IPv6 into the request headers
func (m IPv6Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	rand.Seed(time.Now().UnixNano())
	randomIPv6 := fmt.Sprintf("2001:%x:%x:%x:%x:%x:%x:%x", rand.Intn(65536), rand.Intn(65536), rand.Intn(65536),
		rand.Intn(65536), rand.Intn(65536), rand.Intn(65536), rand.Intn(65536))

	// Set the headers
	r.Header.Set("X-Real-IP", randomIPv6)
	r.Header.Set("X-Forwarded-For", randomIPv6)
	// Debug log (optional)
	//fmt.Println("Injected IPv6:", randomIPv6)

	// Continue request processing
	return next.ServeHTTP(w, r)
}
