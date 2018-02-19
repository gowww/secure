package secure

import "time"

// Options and headers directives.
const (
	FrameDenied     = "DENY"       // The page cannot be displayed in a frame, regardless of the site attempting to do so.
	FrameSameOrigin = "SAMEORIGIN" // The page can only be displayed in a frame on the same origin as the page itself.

	HPKPDefaultMaxAge = 30 * 24 * time.Hour // HPKPDefaultMaxAge provides a default HPKP Max-Age value of 30 days.
	HSTSDefaultMaxAge = 30 * 24 * time.Hour // HSTSDefaultMaxAge provides a default HSTS Max-Age value of 30 days.

	ReferrerPolicyNoReferrer                  = "no-referrer"                     // No referrer information is to be sent.
	ReferrerPolicyNoReferrerWhenDowngrade     = "no-referrer-when-downgrade"      // Send a full URL from a TLS-protected environment settings object to a potentially trustworthy URL, and from clients which are not TLS-protected to any origin.
	ReferrerPolicySameOrigin                  = "same-origin"                     // A full URL, stripped for use as a referrer, is sent when making same-origin requests
	ReferrerPolicyOrigin                      = "origin"                          // Only the ASCII serialization of the origin is sent when making both same-origin and cross-origin requests.
	ReferrerPolicyStrictOrigin                = "strict-origin"                   // Send the ASCII serialization of the origin from a TLS-protected environment settings object to a potentially trustworthy URL, and from non-TLS-protected environment settings objects to any origin.
	ReferrerPolicyOriginWhenCrossOrigin       = "origin-when-cross-origin"        // A full URL, stripped for use as a referrer, is sent when making same-origin requests, and only the ASCII serialization of the origin is sent when making cross-origin requests.
	ReferrerPolicyStrictOriginWhenCrossOrigin = "strict-origin-when-cross-origin" // A full URL, stripped for use as a referrer, is sent when making same-origin requests, and only the ASCII serialization of the origin when making cross-origin requests from a TLS-protected environment settings object to a potentially trustworthy URL, and from non-TLS-protected environment settings objects to any origin.
	ReferrerPolicyUnsafeURL                   = "unsafe-url"                      // A full URL, stripped for use as a referrer, is sent when making both same-origin and cross-origin requests.

	XSSProtectionDisabled = "0"             // Disables XSS filtering.
	XSSProtectionEnabled  = "1"             // Enables XSS filtering (usually default in browsers). If a cross-site scripting attack is detected, the browser will sanitize the page (remove the unsafe parts).
	XSSProtectionBlock    = "1; mode=block" // Enables XSS filtering. Rather than sanitizing the page, the browser will prevent rendering of the page if an attack is detected.
)

// FrameAllow returns a Frame directive.
//
// It allows the page to be displayed in a frame on the specified origin.
func FrameAllow(uri string) string {
	return "ALLOW-FROM " + uri
}

// XSSProtectionReport returns an XSSProtection directive.
//
// It Enables XSS filtering (Chromium only).
// If a cross-site scripting attack is detected, the browser will sanitize the page and report the violation.
// This uses the functionality of the CSP report-uri directive to send a report.
func XSSProtectionReport(uri string) string {
	return "1; report=" + uri
}
