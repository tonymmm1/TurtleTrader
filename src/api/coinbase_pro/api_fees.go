//Coinbase Pro API : Fees
package coinbase_pro

import (
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

func Get_fees() []byte {
    path := "/fees"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
