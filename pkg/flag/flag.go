package flag

import (
	"flag"
)

var (
	Active = ""
)

func ConfigFlag() {
	flag.StringVar(&Active, "active", "", "active profile")
	flag.Parse()
}
