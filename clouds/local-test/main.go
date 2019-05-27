package main

import (
	"net/http"

	. "github.com/carlpett/blackfn_exporter/pkg/config"
	"github.com/carlpett/blackfn_exporter/pkg/probe"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { probe.RunProbe(w, r, Config, Logger) })
	http.ListenAndServe(":8080", nil)
}
