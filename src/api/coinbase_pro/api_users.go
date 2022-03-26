//Coinbase Pro API : Users
package coinbase_pro

import (
    "fmt"
    "os"
)

type Transfer_limit_currency struct {
    Max float64 `json:"max"`
    Remaining float64 `json:"remaining"`
}

type Transfer_limit struct {
    Usd Transfer_limit_currency `json:"USD"`
    Gbp Transfer_limit_currency `json:"GBP"`
    Eur Transfer_limit_currency `json:"EUR"`
    Btc Transfer_limit_currency `json:"BTC"`
    Eth Transfer_limit_currency `json:"ETH"`
}

type User struct {
    Transfer_limits struct {
        Buy Transfer_limit `json:"buy"`
        Sell Transfer_limit `json:"sell"`
        Exchange_withdraw Transfer_limit `json:"exchange_withdraw"`
        Ach Transfer_limit `json:"ach"`
        Ach_no_balance Transfer_limit `json:"transfer_limit"`
        Credit_debit_card Transfer_limit `json:"credit_debit_card"`
        Secure3d_buy Transfer_limit `json:"secure3d_buy"`
        Paypal_buy Transfer_limit `json:"transfer_limit"`
        Paypal_withdrawal Transfer_limit `json:"paypal_withdrawal"`
        Ideal_deposit Transfer_limit `json:"ideal_deposit"`
        Sofort_deposit Transfer_limit `json:"sofort_deposit"`
        Instant_ach_withdrawal Transfer_limit `json:"instant_ach_withdrawal"`
    } `json:"transfer_limits"`
    Limit_currency string `json:"limit_currency"`
}

/*  Users
*       Get user exchange limits    (GET)
*/

func Get_user_exchange_limits(user_id string) []byte {
    path := "/users/" + user_id + "/exchange-limits"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
