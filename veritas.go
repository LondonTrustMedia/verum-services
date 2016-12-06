// written by London Trust Media
// released under the MIT license
package main

import (
	"fmt"
	"log"

	"github.com/Verum/veritas/lib"
	docopt "github.com/docopt/docopt-go"
)

func main() {
	version := lib.SemVer
	usage := `veritas.
Usage:
	veritas initdb [--conf <filename>] [--quiet]
	veritas upgradedb [--conf <filename>] [--quiet]
	veritas genpasswd [--conf <filename>] [--quiet]
	veritas mkcerts [--conf <filename>] [--quiet]
	veritas run [--conf <filename>] [--quiet]
	veritas -h | --help
	veritas --version
	veritas --license
Options:
	--conf <filename>  Configuration file to use [default: services.yaml].
	--quiet            Don't show startup/shutdown lines.
	-h --help          Show this screen.
	--version          Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, version, false)

	configfile := arguments["--conf"].(string)
	config, err := lib.LoadConfig(configfile)
	if err != nil {
		log.Fatal("Config file did not load successfully:", err.Error())
	}

	if arguments["--license"].(bool) {
		fmt.Println(lib.Copyright)
	}
}
