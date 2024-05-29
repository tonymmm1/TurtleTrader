//Config parser
package config

import (
    "fmt"
    "os"

    "github.com/BurntSushi/toml"
)

var Debug bool
var CBP Coinbase_pro
var Refresh float32 = 1.0 //default 1.0

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
}
