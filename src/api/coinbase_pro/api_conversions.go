//Coinbase Pro : Conversions
package coinbase_pro

import (
    "fmt"
    "os"
)

/*  Conversions
*       Convert currency    (POST)
*       Get a conversion    (GET)
*/

type Convert struct { //Convert Currency/Get a conversion
    Id string `json:"id"`
    Amount string `json:"amount"`
    From_account_id string `json:"from_account_id"`
    To_account_id string `json:"to_account_id"`
    From string `json:"from"`
    To string `json:"to"`
}

func convert_currency(profile_id string, from string, to string, amount string, nonce string) []byte { //Converts funds from currency to currency
    path := "/conversions"

    status, response := rest_post_convert(path, profile_id, from, to, amount, nonce)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_conversion(conversion_id string, profile_id string) []byte {
    path := "/conversion/" + conversion_id

    status, response := rest_get_convert(path, profile_id)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
