package csrf

import (
	"crypto/subtle"
	"encoding/base64"
	"net/url"
	"time"

	cfg "github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/lib/strings"
	"github.com/valyala/fasthttp"
)

const (
	TokenLength     = 32
	TokenField      = "csrfToken"
	TokenCookieName = "pnk_csrf"
)

var (
	safeMethods = []string{
		fasthttp.MethodGet,
		fasthttp.MethodHead,
		fasthttp.MethodOptions,
		fasthttp.MethodTrace,
	}
)

type tokenCookie struct {
	Timestamp      int64
	SessionVersion string
	Token          []byte
}

// Protect adds csrf cookie to response
func Protect(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// HTTP methods not defined as idempotent ("safe") under RFC7231 require
		// inspection.
		if !strings.Contains(safeMethods, string(ctx.Method())) {
			decoded := tokenCookie{}
			cookie := ctx.Request.Header.Cookie(TokenCookieName)

			err := cfg.SecureVault.Decode(TokenCookieName, string(cookie), &decoded)
			if err != nil || len(decoded.Token) != TokenLength {
				deny(ctx)
				return
			}

			// Enforce an origin check for HTTPS connections. As per the Django CSRF
			// implementation (https://goo.gl/vKA7GE) the Referer header is almost
			// always present for same-domain HTTP requests.
			if string(ctx.Request.URI().Scheme()) == "https" {
				// Fetch the Referer value. Call the error handler if it's empty or
				// otherwise fails to parse.
				referer, err := url.Parse(string(ctx.Request.Header.Referer()))
				if err != nil || referer.String() == "" {
					deny(ctx)
					return
				}

				requestURL, err := url.Parse(string(ctx.RequestURI()))
				if err != nil || requestURL.String() == "" {
					deny(ctx)
					return
				}

				if !sameOrigin(requestURL, referer) {
					deny(ctx)
					return
				}
			}

			// If the token returned from the session store is nil for non-idempotent
			// ("unsafe") methods, call the error handler.
			if decoded.Token == nil {
				deny(ctx)
				return
			}

			xscrfToken := string(ctx.Request.Header.Peek("X-CSRF-Token"))
			issued, err := base64.StdEncoding.DecodeString(xscrfToken)
			if err != nil {
				deny(ctx)
				return
			}

			requestToken := unmask(issued)

			if !compareTokens(requestToken, decoded.Token) {
				deny(ctx)
				return
			}
		}

		ctx.Response.Header.Add("Vary", "Cookie")
		next(ctx)
	}
}

func deny(ctx *fasthttp.RequestCtx) {
	responses.ClearCookie(ctx, TokenCookieName)
	ctx.SetStatusCode(fasthttp.StatusForbidden)
}

// InjectToken generates new token for template and sets its masked version to cookie
func InjectToken(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := generateToken()

		value, err := cfg.SecureVault.Encode(TokenCookieName, tokenCookie{
			Timestamp:      time.Now().Unix(),
			SessionVersion: cfg.SessionVersion,
			Token:          token,
		})

		if err == nil {
			responses.SetRootCookie(ctx, TokenCookieName, value, cfg.SecondsRarely)
		}

		ctx.SetUserValue(TokenCookieName, mask(token))
		next(ctx)
	}
}

func mask(realToken []byte) string {
	otp := []byte(strings.GenerateRandomString(TokenLength))

	// XOR the OTP with the real token to generate a masked token. Append the
	// OTP to the front of the masked token to allow unmasking in the subsequent
	// request.
	return base64.StdEncoding.EncodeToString(append(otp, xorToken(otp, realToken)...))
}

// unmask splits the issued token (one-time-pad + masked token) and returns the
// unmasked request token for comparison.
func unmask(issued []byte) []byte {
	// Issued tokens are always masked and combined with the pad.
	if len(issued) != TokenLength*2 {
		return nil
	}

	// We now know the length of the byte slice.
	otp := issued[TokenLength:]
	masked := issued[:TokenLength]

	// Unmask the token by XOR'ing it against the OTP used to mask it.
	return xorToken(otp, masked)
}

// xorToken XORs tokens ([]byte) to provide unique-per-request CSRF tokens. It
// will return a masked token if the base token is XOR'ed with a one-time-pad.
// An unmasked token will be returned if a masked token is XOR'ed with the
// one-time-pad used to mask it.
func xorToken(a, b []byte) []byte {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}

	res := make([]byte, n)

	for i := 0; i < n; i++ {
		res[i] = a[i] ^ b[i]
	}

	return res
}

// compare securely (constant-time) compares the unmasked token from the request
// against the real token from the session.
func compareTokens(a, b []byte) bool {
	// This is required as subtle.ConstantTimeCompare does not check for equal
	// lengths in Go versions prior to 1.3.
	if len(a) != len(b) {
		return false
	}

	return subtle.ConstantTimeCompare(a, b) == 1
}

func sameOrigin(a, b *url.URL) bool {
	return (a.Scheme == b.Scheme && a.Host == b.Host)
}

func generateToken() []byte {
	return []byte(strings.GenerateRandomString(TokenLength))
}
