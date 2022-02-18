//Coinbase Pro API : Fees
package coinbase_pro

import (
    "encoding/json"
    "fmt"
    "os"
)

type Fee struct { //Get fees
    Taker_fee_rate string `json:"taker_fee_rate"`
    Maker_fee_rate string `json:"maker_fee_rate"`
    Usd_volume string `json:"usd_volume"`
}

/*  Fees
*       Get fees    (GET)
*/

func Get_fees() Fee {
    path := "/fees"

    var fees Fee

    response_status, response_body := rest_get(path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &fees); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get fees")
    fmt.Println()
    fmt.Println(fees.Taker_fee_rate)
    fmt.Println(fees.Maker_fee_rate)
    fmt.Println(fees.Usd_volume)
    fmt.Println()

    return fees
}
