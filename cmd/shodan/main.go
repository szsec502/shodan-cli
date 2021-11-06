package main

import (
	"os"

	"github.com/seaung/shodan-cli/pkg/lib"
	"github.com/seaung/shodan-cli/shodan"
)

func main() {
	options := lib.ParseOptions()

	shodan := shodan.New(os.Getenv("SHODAN_API_KEY"))

	if options.Domain != "" {
		shodan.DomainInfo(options.Domain)
	}

}
