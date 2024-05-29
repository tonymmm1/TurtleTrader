package main

import (
	"flag"
	"fmt"
	"os"
	"turtle/src/config"
	//cbp "turtle/src/api/coinbase_pro"
	"turtle/src/ui"
)

type CbpConfig struct { //Coinbase Pro configuration
    Host string
    Key string
    Password string
    Secret string
}

func init() {
	flag.BoolVar(&config.Debug, "d", false, "enable debug mode")
	flag.BoolVar(&config.Debug, "debug", false, "enable debug mode")
}

func main() {
	help := flag.Bool("help", false, "Display help information")
	flag.BoolVar(help, "h", false, "Display help information")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "TurtleTrader: crypto-currency trading bot\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n", os.Args[0])

		// Manually format each flag
		fmt.Fprintf(flag.CommandLine.Output(), "  -d        %s\n", "Enable debug mode")
		fmt.Fprintf(flag.CommandLine.Output(), "  -debug    %s\n", "Enable debug mode")
		fmt.Fprintf(flag.CommandLine.Output(), "  -h        %s\n", "Display help information")
		fmt.Fprintf(flag.CommandLine.Output(), "  -help     %s\n", "Display help information")
	}

	//Parse flags
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

    fmt.Println("TurtleTrader")
    fmt.Println()

	if config.Debug {
		fmt.Println("Debug mode enabled")
	}

    fmt.Println("Loading api.toml")
    config.Load()
    
    fmt.Println()

	ui.Init()
}
