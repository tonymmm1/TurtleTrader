package main

import (
        "fmt"
        "os"

        "github.com/BurntSushi/toml"
)

var cbpKey cbpConfig //global struct that stores Coinbase Pro credentials

func main() {
    f := "api.toml" //default configuration file
    if _, err := os.Stat(f); err != nil { //check configuration file exists
        fmt.Println("ERROR " + f + " does not exist")
        os.Exit(1)
    }

    if _, err := toml.DecodeFile(f, &cbpKey); err != nil { //decode TOML file
        fmt.Println("ERROR decoding toml configuration")
        os.Exit(1)
    }

    //add check for missing config elements
   
    rest_handler() //rest handle function
}
