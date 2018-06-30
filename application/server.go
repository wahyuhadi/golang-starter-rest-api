/**
	Rahmat wahyu hadi
**/

package application

import (
	"./middlewares"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

// NewServer ...
func NewServer() *negroni.Negroni {
	// Untuk Global Middleware
	server := negroni.New()
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	server.Use(middlewares.CORSMiddleware())

	server.UseHandler(NewRouter())

	return server
}
