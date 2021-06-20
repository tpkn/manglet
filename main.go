package main

import (
	"flag"
	"fmt"
	"log"
	
	"manglet/manglet"
	"manglet/models"
	"manglet/utils"
)

var env = "development"
var version = "0.0.0"
var description = fmt.Sprintf(`Manglet, v%v | https://tpkn.me
Tool for quick mangle sensitive data in textish files .
It replace all letters by 'x', all digits by '0' and all symbols except '-,' by '-'.
There are no other mangle scenarios yet.`, version)

func main() {
	cli := models.CLI{}
	flag.StringVar(&cli.Input, "i", "", "Input file path")
	flag.StringVar(&cli.Output, "o", "", "Output file path. If empty, then input file path with suffix '_manglet' will be used instead")
	flag.StringVar(&cli.MangleGroups, "m", "ld", "Mangle groups ('l' = letters; 'd' = digits; 's' = symbols)")
	flag.StringVar(&cli.Skip, "skip", "", "Rows to skip, separated by commas")
	flag.StringVar(&cli.About, "about", "", description)
	flag.BoolVar(&cli.Verbose, "v", false, "Verbose output")
	flag.Parse()
	
	if env == "development" {
		// cli.Input = "./_/test_500k.csv"
		cli.Input = "./_/test_short.txt"
		cli.Output = "X:/"
		cli.MangleGroups = "dl"
		cli.Skip = "1"
	}
	
	if !utils.FileExists(cli.Input) {
		log.Fatalf("Input file does not exist")
	}
	
	// Start process
	mangled_file, err := manglet.Run(&cli)
	if err != nil {
		panic(err)
	}
	
	// Print mangled file path to STDOUT
	if cli.Verbose {
		fmt.Print(mangled_file)
	}
}
