//Config parser
package config

import (
    "fmt"
    "os"

    "github.com/BurntSushi/toml"
)

var CBP Coinbase_pro

func Load() {
    file := "api.toml" //default configuration file
    if _, err := os.Stat(file); err != nil { //check configuration file exists
        fmt.Println("ERROR " + file + " does not exist")
        os.Exit(1)
    }

    if _, err := toml.DecodeFile(file, &CBP); err != nil { //decode TOML file
        fmt.Println("ERROR decoding toml configuration")
        os.Exit(1)
    }

    //debug 
    fmt.Println(CBP)
}
