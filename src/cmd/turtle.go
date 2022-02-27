package main

import (
        "fmt"
        
        "turtle/src/config"
        //cbp "turtle/src/api/coinbase_pro"
)

type CbpConfig struct { //Coinbase Pro configuration
    Host string
    Key string
    Password string
    Secret string
}


func usage() {
    usage := 
    `TurtleTrader: crypto-currency trading bot
Usage: turtle [options] ...
    -h,--help       Show this help message
    -v,--verbose    Show verbose output`
    fmt.Printf("%s\n\n", usage)
}

func main() {
    fmt.Println("TurtleTrader")
    fmt.Println()

    fmt.Println("Loading api.toml")
    config.Load()
    
    fmt.Println()

    usage()

    fmt.Println()
}
