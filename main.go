package main

import (
	"net/http"

	"golang.org/x/crypto/acme/autocert"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/meyskens/readmyage-api/config"
)

var conf config.Config

func main() {
	conf = config.GetConfig()

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", serveRoot)
	e.GET("/lookup/isbn", servelLookupISBN)

	if conf.AutoTLS {
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(conf.Hostname)
		e.AutoTLSManager.Cache = autocert.DirCache(conf.CertCache)
		e.Logger.Fatal(e.StartAutoTLS(conf.Bind))
	} else {
		e.Logger.Fatal(e.Start(conf.Bind))
	}
}

func serveRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! This code is rated ages 3-999")
}

func servelLookupISBN(c echo.Context) error {
	isbn := c.QueryParam("isbn")

	done := make(chan bool)
	in := make(chan LookUpResult)
	response := LookUpResponse{
		Results: []LookUpResult{},
	}

	go lookUpBibliotheek(isbn, done, in)

	for i := 0; i < 1; i++ {
		succeed := <-done
		if succeed {
			response.Results = append(response.Results, <-in)
		}
	}

	return c.JSON(http.StatusOK, response)
}
