package main

import (
	"net/http"

	"github.com/prest/cmd"
	"github.com/prest/config"
	"google.golang.org/appengine"
)

func main() {
	config.Load()
	http.Handle("/", cmd.MakeHandler())
	appengine.Main()
}
