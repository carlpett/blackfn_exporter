package function

import (
	"net/http"

	. "github.com/carlpett/blackfn_exporter/pkg/config"
	"github.com/carlpett/blackfn_exporter/pkg/probe"
)

func BlackboxProbe(w http.ResponseWriter, r *http.Request) {
	probe.RunProbe(w, r, Config, Logger)
}
