package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"

	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/meyskens/readmyage-api/config"
)

var conf config.Config
var db *sql.DB

func main() {
	conf = config.GetConfig()

	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

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
	if isbn == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "please specify an ISBN number"})
	}

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

	go logQuery(isbn, len(response.Results))

	return c.JSON(http.StatusOK, response)
}

func logQuery(isbn string, results int) { // keeping a database of ISBNs during testing to test new data sets
	db.Exec("INSERT INTO lookup (isbn, results) VALUES ($1, $2)", isbn, results)
}
