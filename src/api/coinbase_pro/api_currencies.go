//Coinbase Pro API : Currencies
package coinbase_pro

import (
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

func Get_all_currencies() []byte {
    path := "/currencies"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_currency(currency_id string) []byte {
    path := "/currencies/" + currency_id

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
