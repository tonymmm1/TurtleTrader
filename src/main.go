package main

import (
        "fmt"
        "os"

        "github.com/BurntSushi/toml"
)

func main() {
    f := "api.toml" //default configuration file
    if _, err := os.Stat(f); err != nil { //check configuration file exists
        fmt.Println("ERROR " + f + " does not exist")
        os.Exit(1)
    }

    var api_struct apiConfig //struct that stores API key information
    if _, err := toml.DecodeFile(f, &api_struct); err != nil { //decode TOML file
        fmt.Println("ERROR decoding toml configuration")
        os.Exit(1)
    }

    //add check for missing config elements
   
    rest_handler(api_struct) //rest handle function
}
