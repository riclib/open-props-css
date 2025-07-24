package uicss

import (
	_ "embed"
	"net/http"
	"strings"
)

//go:embed dashboard.css
var dashboardCSS string

//go:embed dashboard.readable.css
var dashboardReadableCSS string

//go:embed stylebook.html
var stylebookHTML string

// CSS returns the raw CSS as a string
func CSS() string {
	return dashboardCSS
}

// ReadableCSS returns the non-minified CSS as a string
func ReadableCSS() string {
	return dashboardReadableCSS
}

// CSSHandler returns an http.Handler that serves the dashboard.css file
func CSSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.Header().Set("Cache-Control", "public, max-age=3600")
		w.Write([]byte(dashboardCSS))
	})
}

// ReadableCSSHandler returns an http.Handler that serves the readable CSS file
func ReadableCSSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.Header().Set("Cache-Control", "public, max-age=3600")
		w.Write([]byte(dashboardReadableCSS))
	})
}

// StylebookHandler returns an http.Handler that serves the interactive stylebook
func StylebookHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		
		// Inject the CSS inline into the stylebook HTML
		html := strings.Replace(stylebookHTML, "<!-- CSS_PLACEHOLDER -->", "<style>"+dashboardCSS+"</style>", 1)
		w.Write([]byte(html))
	})
}