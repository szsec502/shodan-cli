package lib

import (
	"flag"

	"github.com/seaung/shodan-cli/pkg/ascii"
)

type Options struct {
	Domain string
	Host   string
	Info   bool
}

func ParseOptions() *Options {
	options := &Options{}

	flag.StringVar(&options.Domain, "domain", "", "target domain.")
	flag.BoolVar(&options.Info, "info", false, "show api info")
	flag.StringVar(&options.Host, "host", "", "target host addresses.")

	flag.Parse()

	ascii.ShowBanner()

	return options
}
