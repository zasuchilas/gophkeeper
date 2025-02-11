package main

import "github.com/zasuchilas/gophkeeper/internal/server/app"

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	a := app.New(buildVersion, buildDate, buildCommit)
	a.Run()
}
