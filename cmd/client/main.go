package main

import "github.com/zasuchilas/gophkeeper/internal/client/app"

var (
	buildVersion = ""
	buildDate    = ""
	buildCommit  = ""
)

func main() {
	a := app.New(buildVersion, buildDate, buildCommit)
	a.Run()
}
