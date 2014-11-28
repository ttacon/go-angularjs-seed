package main

import (
	"flag"
	"github.com/bmizerany/pat"
	"github.com/yext/glog"
	"net/http"
)

var (
	port     = flag.String("p", "1809", "port to run on")
	prodMode = flag.Bool("prod", false, "run in prod mode")
)

func main() {
	flag.Parse()

	m := pat.New()

	http.Handle("/api/", m)
	http.Handle("/", http.FileServer(http.Dir("static/")))
	glog.Fatal(http.ListenAndServe(":"+*port, nil))
}
