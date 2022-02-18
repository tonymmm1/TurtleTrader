package main

import (
        "fmt"
        "os"

        "crypto-bot/src/api/coinbase_pro"

        "github.com/BurntSushi/toml"
)

type cbpConfig struct { //Coinbase Pro configuration
    Host string
    Key string
    Password string
    Secret string
}

var CbpKey cbpConfig

func main() {
    f := "api.toml" //default configuration file
    if _, err := os.Stat(f); err != nil { //check configuration file exists
        fmt.Println("ERROR " + f + " does not exist")
        os.Exit(1)
    }

    if _, err := toml.DecodeFile(f, &CbpKey); err != nil { //decode TOML file
        fmt.Println("ERROR decoding toml configuration")
        os.Exit(1)
    }

    //add check for missing config elements
   
    rest_handler() //rest handle function
}
