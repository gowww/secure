// Package secure provides security utilities, CSP, HPKP, HSTS and other security wins.
package secure

import (
	"fmt"
	"net/http"
	"time"
)

// HSTSPreloadMinAge is the lowest max age usable with HSTS preload. See https://hstspreload.appspot.com.
const HSTSPreloadMinAge = 10886400

// A handler provides a security handler.
type handler struct {
	next    http.Handler
	options *Options
}

// Options represents security options.
type Options struct {
	AllowedHosts   []string     // AllowedHosts indicates which fully qualified domain names are allowed to point to this server. If none are set, all are allowed.
	CSP            string       // CSP contains Content Security Policy. See http://www.w3.org/TR/CSP and https://developer.mozilla.org/en-US/docs/Web/Security/CSP/Using_Content_Security_Policy.
	Frame          string       // FrameAllowed indicates whether or not a browser should be allowed to render a page in a frame, iframe or object. Default is FrameSameOrigin.
	HPKP           *HPKPOptions // HPKP contains the HTTP Public Key Pinning options.
	HSTS           *HSTSOptions // HPKP contains the HTTP Strict Transport Security options.
	ReferrerPolicy string       // ReferrerPolicy contains Referrer Policy. See https://www.w3.org/TR/referrer-policy.
	XSSProtection  string       // XSSProtection can stop pages from loading when browser detects an XSS attack. Default is XSSProtectionBlock.
	ForceSSL       bool         // ForceSSL indicates whether an insecure request must be redirected to the secure protocol.
	EnvDevelopment bool         // EnvDevelopment can be used during development to defuse AllowedHosts, HPKP, HSTS and ForceSSL options.
}

// HPKPOptions represents HTTP Public Key Pinning options.
// See RFC 7469 and https://developer.mozilla.org/en-US/docs/Web/Security/Public_Key_Pinning.
type HPKPOptions struct {
	Keys              []string      // Keys contains the Base64 encoded Subject Public Key Information (SPKI) fingerprints. This field is required.
	MaxAge            time.Duration // MaxAge indicates how long the browser should remember that this site is only to be accessed using one of the pinned keys. This field is required.
	IncludeSubdomains bool          // IncludeSubdomains indicates whether HPKP applies to all of the site's subdomains as well.
	ReportURI         string        // ReportURI is the URL at which validation failures are reported to.
}

// HSTSOptions represents HTTP Strict Transport Security options.
// See RFC 6797 and https://developer.mozilla.org/en-US/docs/Web/Security/HTTP_strict_transport_security.
type HSTSOptions struct {
	MaxAge            time.Duration // MaxAge indicates how long the browser should remember that this site is only to be accessed using HTTPS. This field is required.
	IncludeSubdomains bool          // IncludeSubdomains indicates whether HSTS applies to all of the site's subdomains as well.
	Preload           bool          // Preload indicates whether the browsers must use a secure connection. It's not a standard. See https://hstspreload.appspot.com.
}

// Handle returns a Handler wrapping another http.Handler.
func Handle(h http.Handler, o *Options) http.Handler {
	// Validate options (with required fields) from the beginning.
	if o != nil {
		if o.Frame == "" {
			o.Frame = FrameSameOrigin
		}
		if o.HPKP != nil {
			hpkpHeader(o)
		}
		if o.HSTS != nil {
			hstsHeader(o)
		}
		if o.XSSProtection == "" {
			o.Frame = XSSProtectionBlock
		}
	}
	return &handler{h, o}
}

// HandleFunc returns a Handler wrapping an http.HandlerFunc.
func HandleFunc(f http.HandlerFunc, o *Options) http.Handler {
	return Handle(f, o)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.options != nil {
		if !h.options.EnvDevelopment {
			if len(h.options.AllowedHosts) > 0 {
				for _, host := range h.options.AllowedHosts {
					if host == r.URL.Host {
						goto SSLOptions
					}
				}
				http.NotFound(w, r)
				return
			}
		SSLOptions:
			isSSL := r.URL.Scheme == "https" || r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https"
			if !isSSL && h.options.ForceSSL {
				r.URL.Scheme = "https"
				http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
				return
			}
			if isSSL && h.options.HPKP != nil {
				w.Header().Set("Public-Key-Pins", hpkpHeader(h.options))
			}
			if h.options.HSTS != nil {
				w.Header().Set("Strict-Transport-Security", hstsHeader(h.options))
			}
		}
		if h.options.CSP != "" {
			w.Header().Set("Content-Security-Policy", h.options.CSP)
			w.Header().Set("X-Content-Security-Policy", h.options.CSP)
			w.Header().Set("X-WebKit-CSP", h.options.CSP)
		}
		if h.options.ReferrerPolicy != "" {
			w.Header().Add("Referrer-Policy", h.options.ReferrerPolicy)
		}
	}

	// Good practice headers.
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", string(h.options.Frame))
	w.Header().Set("X-XSS-Protection", h.options.XSSProtection)

	h.next.ServeHTTP(w, r)
}

func hpkpHeader(o *Options) (v string) {
	if len(o.HPKP.Keys) == 0 {
		panic("secure: at least one key must be set when using HPKP")
	}
	if o.HPKP.MaxAge == 0 {
		panic("secure: max age must be set when using HPKP")
	}

	for _, key := range o.HPKP.Keys {
		if v != "" {
			v += "; "
		}
		v += fmt.Sprintf("pin-sha256=%q", key)
	}
	v += fmt.Sprintf("; %.f", o.HPKP.MaxAge.Seconds())
	if o.HPKP.IncludeSubdomains {
		v += "; includeSubdomains"
	}
	if o.HPKP.ReportURI != "" {
		v += fmt.Sprintf("; report-uri=%q", o.HPKP.ReportURI)
	}
	return
}

func hstsHeader(o *Options) (v string) {
	if !o.ForceSSL {
		panic("secure: ForceSSL must be true when using HSTS")
	}
	if o.HSTS.MaxAge == 0 {
		panic("secure: max age must be set when using HSTS")
	}
	if o.HSTS.Preload {
		if o.HSTS.MaxAge < HSTSPreloadMinAge {
			panic("secure: max age must be at least 18 weeks when using HSTS preload")
		}
		if !o.HSTS.IncludeSubdomains {
			panic("secure: subdomains must be included when using HSTS preload")
		}
	}

	v += fmt.Sprintf("%.f", o.HSTS.MaxAge.Seconds())
	if o.HSTS.IncludeSubdomains {
		v += "; includeSubdomains"
	}
	if o.HSTS.Preload {
		v += "; preload"
	}
	return
}
