//Coinbase Pro API : Coinbase price oracle
package coinbase_pro

import (
    "fmt"
    "os"
)

type Price struct { //Get signed prices
    Timestamp string `json:"timestamp"`
//    Messages []string `json:"messages"`
//    Signatures []string `json:"signatures"`
    Prices map[string] interface {} `json:"prices"`
}

/*  Coinbase price oracle
*       Get signed prices   (GET)
*/

func Get_signed_prices() []byte {
    path := "/oracle"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
