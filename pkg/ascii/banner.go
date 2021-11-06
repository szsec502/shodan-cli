package ascii

import (
	"github.com/projectdiscovery/gologger"
)

const banner = `
        _,--',   _._.--._____
 .--.--';_'-.', ";_      _.,-'
.'--'.  _.'    {''-;_ .-.>.'
      '-:_      )  / '' '=.
        ) >     {_/,     /~)    v0.0.0
        |/               '^ .'
`

const VERSION = `v0.0.0`

func ShowBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tgithub.com/seaung\n\n")
	gologger.Print().Msgf("A shodan command cli tools.\n")
}
