package lib

import (
	"flag"
	"github.com/seaung/shodan-cli/pkg/ascii"
)

type Options struct {
	Domain string
	Host   string
	Info   string
}

func ParseOptions() *Options {
	options := &Options{}

	flag.StringVar(&options.Domain, "domain", "", "target domain.")
	flag.StringVar(&options.Info, "info", "", "show api info.")
	flag.StringVar(&options.Host, "host", "", "target host addresses.")

	flag.Parse()

	ascii.ShowBanner()

	return options
}
