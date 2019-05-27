package main

import (
	"net/http"

	. "github.com/carlpett/blackfn_exporter/pkg/config"
	"github.com/carlpett/blackfn_exporter/pkg/probe"

	"github.com/imdario/gluo"
)

func BlackboxProbe(w http.ResponseWriter, r *http.Request) {
	probe.RunProbe(w, r, Config, Logger)
}

func main() {
	gluo.ListenAndServe(":8080", http.HandlerFunc(BlackboxProbe))
}
