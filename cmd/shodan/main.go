package main

import (
	"fmt"
	"os"

	"github.com/seaung/shodan-cli/pkg/lib"
	"github.com/seaung/shodan-cli/shodan"
)

func main() {
	options := lib.ParseOptions()

	shodan := shodan.New(os.Getenv("SHODAN_API_KEY"))

	if !options.Info {
		apiInfo, err := shodan.APIInfo()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(apiInfo)
	}

	if options.Domain != "" {
		domain, err := shodan.DomainInfo(options.Domain)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(domain)
	}

	if options.Host != "" {
		host, err := shodan.HostSearch(options.Host)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(host)
	}
}
