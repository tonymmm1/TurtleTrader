//Coinbase Pro API : Currencies
package coinbase_pro

import (
    "encoding/json"
    "fmt"
    "os"
)

type Currency struct { //Get all known currency/Get a currency
    Id string `json:"id"`
    Name string `json:"name"`
    Min_size string `json:"min_size"`
    Status string `json:"status"`
    Message string `json:"message"`
    Max_precision string `json:"max_precision"`
    Convertible_to []string `json:"convertible_to"`
    Details struct {
        Type string `json:"type"`
        Symbol string `json:"symbol"`
        Network_confirmations int32 `json:"network_confirmations"`
        Sort_order int32 `json:"sort_order"`
        Crypto_address_link string `json:"crypto_address_link"`
        Crypto_transaction_link string `json:"crypto_transaction_link"`
        Push_payment_methods []string `json:"push_payment_methods"`
        Group_types []string `json:"group_types"`
        Display_name string `json:"display_name"`
        Processing_time_seconds float64 `json:"processing_time_seconds"`
        Min_withdrawal_amount float64 `json:"min_withdrawal_amount"`
        Max_withdrawal_amount float64 `json:"max_withdrawal_amount"`
    } `json:"details"`
}

/*  Currencies
*       Get all known currencies    (GET)
*       Get a currency              (GET)
*/

func Get_all_currencies() []Currency {
    path := "/currencies"

    var currencies []Currency

    response_status, response_body := rest_get(path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &currencies); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get all known currencies")
    fmt.Println()
    for currency := range currencies {
        fmt.Println("currencies[", currency, "]")
        fmt.Println(currencies[currency].Id)
        fmt.Println(currencies[currency].Name)
        fmt.Println(currencies[currency].Min_size)
        fmt.Println(currencies[currency].Status)
        fmt.Println(currencies[currency].Message)
        fmt.Println(currencies[currency].Max_precision)
        for convert := range currencies[currency].Convertible_to {
            fmt.Println("currencies[", currency, "].Convertible_to[", convert, "]")
            fmt.Println(currencies[currency].Convertible_to[convert])
        }
        fmt.Println(currencies[currency].Details.Type)
        fmt.Println(currencies[currency].Details.Symbol)
        fmt.Println(currencies[currency].Details.Network_confirmations)
        fmt.Println(currencies[currency].Details.Sort_order)
        fmt.Println(currencies[currency].Details.Crypto_address_link)
        fmt.Println(currencies[currency].Details.Crypto_transaction_link)
        for payment := range currencies[currency].Details.Push_payment_methods {
            fmt.Println("currencies[", currency, "].Details.Push_payment_methods[", payment, "]")
            fmt.Println(currencies[currency].Details.Push_payment_methods[payment])
        }
        for group := range currencies[currency].Details.Group_types {
            fmt.Println("currencies[", currency, "].Details.Group_types[", group, "]")
            fmt.Println(currencies[currency].Details.Group_types[group])
        }
        fmt.Println(currencies[currency].Details.Display_name)
        fmt.Println(currencies[currency].Details.Processing_time_seconds)
        fmt.Println(currencies[currency].Details.Min_withdrawal_amount)
        fmt.Println(currencies[currency].Details.Max_withdrawal_amount)
        fmt.Println()
    }

    return currencies
}

func Get_currency(currency_id string) Currency {
    path := "/currencies/" + currency_id

    var currency Currency

    response_status, response_body := rest_get(path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &currency); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get a currency")
    fmt.Println()
    fmt.Println("currency:", currency_id)
    fmt.Println(currency.Id)
    fmt.Println(currency.Name)
    fmt.Println(currency.Min_size)
    fmt.Println(currency.Status)
    fmt.Println(currency.Message)
    fmt.Println(currency.Max_precision)
    for convert := range currency.Convertible_to {
        fmt.Println("currency.Convertible_to[", convert, "]")
        fmt.Println(currency.Convertible_to[convert])
    }
    fmt.Println(currency.Details.Type)
    fmt.Println(currency.Details.Symbol)
    fmt.Println(currency.Details.Network_confirmations)
    fmt.Println(currency.Details.Sort_order)
    fmt.Println(currency.Details.Crypto_address_link)
    fmt.Println(currency.Details.Crypto_transaction_link)
    for payment := range currency.Details.Push_payment_methods {
        fmt.Println("currency.Details.Push_payment_methods[", payment, "]")
        fmt.Println(currency.Details.Push_payment_methods[payment])
    }
    for group := range currency.Details.Group_types {
        fmt.Println("currency.Details.Group_types[", group, "]")
        fmt.Println(currency.Details.Group_types[group])
    }
    fmt.Println(currency.Details.Display_name)
    fmt.Println(currency.Details.Processing_time_seconds)
    fmt.Println(currency.Details.Min_withdrawal_amount)
    fmt.Println(currency.Details.Max_withdrawal_amount)
    fmt.Println()

    return currency
}
