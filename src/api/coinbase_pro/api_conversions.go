//Coinbase Pro : Conversions
package coinbase_pro

import (
    "encoding/json"
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

func convert_currency(profile_id string, from string, to string, amount string, nonce string) Convert { //Converts funds from currency to currency
    path := "/conversions"

    var convert Convert

    response_status, response_body := rest_post_convert(path, profile_id, from, to, amount, nonce)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &convert); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Convert currency")
    fmt.Println()
    fmt.Println(convert.Id)
    fmt.Println(convert.Amount)
    fmt.Println(convert.From_account_id)
    fmt.Println(convert.To_account_id)
    fmt.Println(convert.From)
    fmt.Println(convert.To)

    return convert
}

func Get_conversion(conversion_id string, profile_id string) Convert{
    path := "/conversion/" + conversion_id

    var convert Convert

    response_status, response_body := rest_get_convert(path, profile_id)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &convert); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get a conversion")
    fmt.Println()
    fmt.Println(convert.Id)
    fmt.Println(convert.Amount)
    fmt.Println(convert.From_account_id)
    fmt.Println(convert.To_account_id)
    fmt.Println(convert.From)
    fmt.Println(convert.To)

    return convert
}
