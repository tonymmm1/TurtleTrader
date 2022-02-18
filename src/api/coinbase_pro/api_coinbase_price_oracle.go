//Coinbase Pro API : Coinbase price oracle
package coinbase_pro

import (
    "encoding/json"
    "fmt"
    "os"
)

type Price struct { //Get signed prices
    Timestamp string `json:"timestamp"`
    Messages []string `json:"messages"`
    Signatures []string `json:"signatures"`
    Prices map[string] interface {} `json:"prices"`
}

/*  Coinbase price oracle
*       Get signed prices   (GET)
*/

func Get_signed_prices() Price {
    path := "/oracle"

    var prices Price

    response_status, response_body := rest_get(path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &prices); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Print("Get signed prices")
    fmt.Println()
    fmt.Println(prices.Timestamp)
    for message := range prices.Messages {
        fmt.Println("prices.Messages[", message, "]")
        fmt.Println(prices.Messages[message])
        fmt.Println()
    }
    fmt.Println()
    for signature := range prices.Signatures {
        fmt.Println("price.Signatures[", signature, "]")
        fmt.Println(prices.Signatures[signature])
        fmt.Println()
    }
    fmt.Println()
    for k, v := range prices.Prices {
        fmt.Println(k, ":", v)
    }
    fmt.Println()

    return prices
}
