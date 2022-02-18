//Coinbase Pro API : Transfers
package coinbase_pro

import (
    "encoding/json"
    "fmt"
    "os"
)

type cbpTransfer struct { //Withdraw/deposit to/from Coinbase/payment
    Id string `json:"id"`
    Amount string `json:"amount"`
    Currency string `json:"currency"`
    Payout_at string `json:"payout_at"`
    Fee string `json:"fee"`
    Subtotal string `json:"subtotal"`
}

type cbpPayment struct { //Get all payment methods
    Id string `json:"id"`
    Type string `json:"type"`
    Name string `json:"name"`
    Currency string `json:"currency"`
    Primary_buy bool `json:"primary_buy"`
    Primary_sell bool `json:"primary_sell"`
    Instant_buy bool `json:"instant_buy"`
    Instant_sell bool `json:"instant_sell"`
    Created_at string `json:"created_at"`
    Updated_at string `json:"updated_at'`
    Resource string `json:"resource"`
    Resource_path string `json:"resource_path'`
    Verified bool `json:"verified"`
    Limits struct {
        Buy []struct {
            Period_in_days int32 `json:"period_in_days"`
            Total struct {
                Amount string `json:"amount"`
                Currency string `json:"currency"`
            } `json:"total"`
            Remaining struct {
                Amount string `json:"amount"`
                Currency string `json:"currency"`
            } `json:"remaining"`
        } `json:"buy"`
        Instant_buy []struct {
            Period_in_days int32 `json:"period_in_days"`
            Total struct {
                Amount string `json:"amount"`
                Currency string `json:"currency"`
            } `json:"total"`
            Remaining struct {
                Amount string `json:"amount"`
                Currency string `json:"currency"`
            } `json:"remaining"`
        } `json:"instant_buy"`
        Sell []struct {
            Period_in_days int32 `json:"period_in_days"`
            Total struct {
                Amount string `json:"amount"`
                Currency string `json:"currency"`
            } `json:"total"`
            Remaining struct {
                Amount string `json:"amount"`
                Currency string `json:"currency"`
            } `json:"remaining"`
        } `json:"sell"`
        Deposit []struct {
            Period_in_days int32 `json:"period_in_days"`
            Total struct {
                Amount string `json:"amount"`
                Currency string `json:"currency"`
            } `json:"total"`
            Remaining struct {
                Amount string `json:"amount"`
                Currency string `json:"currency"`
            } `json:"remaining"`
        } `json:"deposit"`
    } `json:"limits"`
    Allow_buy bool `json:"allow_buy"`
    Allow_sell bool `json:"allow_sell"`
    Allow_deposit bool `json:"allow_deposit"`
    Allow_withdraw bool `json:"allow_withdraw"`
    Fiat_account struct {
        Id string `json:"id"`
        Resource string `json:"resource"`
        Resource_path string `json:"resource_path"`
    } `json:"fiat_account"`
    Crypto_account struct {
        Id string `json:"id"`
        Resource string `json:"resource"`
        Resource_path string `json:"resouce_path"`
    } `json:"crypto_account"`
    Recurring_options []struct {
        Period string `json:"period"`
        Label string `json:"label"`
    } `json:"recurring_options"`
    Available_balance struct {
        Amount string `json:"amount"`
        Currency string `json:"currency"`
        Scale string `json:"scale"`
    } `json:"available_balance"`
    Picker_data struct {
        Symbol string `json:"symbol"`
        Customer_name string `json:"customer_name"`
        Account_name string `json:"account_name"`
        Account_number string `json:"account_number"`
        Account_type string `json:"account_type"`
        Institution_code string `json:"institution_code"`
        Institution_name string `json:"institution_name"`
        Iban string `json:"iban"`
        Swift string `json:"swift"`
        Paypal_email string `json:"paypal_email"`
        Paypal_owner string `json:"paypal_owner"`
        Routing_number string `json:"routing_name"`
        Institution_identifier string `json:"institution_identifier"`
        Bank_name string `json:"bank_name"`
        Branch_name string `json:"branch_name"`
        Icon_url string `json:"icon_url"`
        Balance struct {
            Amount string `json:"amount"`
            Currency string `json:"currency"`
        } `json:"balance"`
    } `json:"Picker_data"`
    Hold_business_days int64 `json:"hold_business_days"`
    Hold_days int64 `json:"hold_days"`
    Verificationmethod string `json:"verificationMethod"`
    Cdvstatus string `json:"cdvStatus"`
}

/*  Transfers
*       Deposit/Withdraw to/from Coinbase/payment   (POST)
*       Get all payment methods                     (GET)
*       Get all transfers                           (GET)
*       Get a single transfer                       (GET)
*       Get fee estimate for crypto withdrawal      (GET)
*/

func cbp_transfer_coinbase_account(profile_id string, amount string, account_id string, currency string) cbpTransfer {
    path := "/deposits/coinbase-account"

    var deposit cbpTransfer

    response_status, response_body := cbp_rest_post_coinbase(path, profile_id, amount, account_id, currency)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &deposit); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Deposit/Withdraw to/from Coinbase account")
    fmt.Println()
    fmt.Println(deposit.Id)
    fmt.Println(deposit.Amount)
    fmt.Println(deposit.Currency)
    fmt.Println(deposit.Payout_at)
    fmt.Println(deposit.Fee)
    fmt.Println(deposit.Subtotal)
    fmt.Println()

    return deposit
}

func cbp_transfer_payment_account(profile_id string, amount string, account_id string, currency string) cbpTransfer {
    path := "/deposits/payment-method"

    var deposit cbpTransfer

    response_status, response_body := cbp_rest_post_payment(path, profile_id, amount, account_id, currency)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &deposit); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Deposit/Withdraw to/from Coinbase account")
    fmt.Println()
    fmt.Println(deposit.Id)
    fmt.Println(deposit.Amount)
    fmt.Println(deposit.Currency)
    fmt.Println(deposit.Payout_at)
    fmt.Println(deposit.Fee)
    fmt.Println(deposit.Subtotal)
    fmt.Println()

    return deposit
}

func cbp_get_all_payments() []cbpPayment {
    path := "/payment-methods"

    var payments []cbpPayment

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &payments); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get all payment methods")
    fmt.Println()
    for payment := range payments {
        fmt.Println("payments[", payment, "]")
        fmt.Println(payments[payment].Id)
        fmt.Println(payments[payment].Type)
        fmt.Println(payments[payment].Name)
        fmt.Println(payments[payment].Currency)
        fmt.Println(payments[payment].Primary_buy)
        fmt.Println(payments[payment].Primary_sell)
        fmt.Println(payments[payment].Instant_buy)
        fmt.Println(payments[payment].Instant_sell)
        fmt.Println(payments[payment].Created_at)
        fmt.Println(payments[payment].Updated_at)
        fmt.Println(payments[payment].Resource)
        fmt.Println(payments[payment].Resource_path)
        for buy := range payments[payment].Limits.Buy {
            fmt.Println("payments[payment].Limits.Buy[", buy, "]")
            fmt.Println("payments[payment].Limits.Buy[", buy, "].Total")
            fmt.Println(payments[payment].Limits.Buy[buy].Period_in_days)
            fmt.Println(payments[payment].Limits.Buy[buy].Total.Amount)
            fmt.Println(payments[payment].Limits.Buy[buy].Total.Currency)
            fmt.Println("payments[payment].Limits.Buy[", buy, "].Remaining")
            fmt.Println(payments[payment].Limits.Buy[buy].Remaining.Amount)
            fmt.Println(payments[payment].Limits.Buy[buy].Remaining.Currency)
        }
        fmt.Println()
        for instant := range payments[payment].Limits.Instant_buy {
            fmt.Println("payments[payment].Limits.Instant_buy[", instant, "]")
            fmt.Println("payments[payment].Limits.Instant_buy[", instant, "].Total")
            fmt.Println(payments[payment].Limits.Instant_buy[instant].Period_in_days)
            fmt.Println(payments[payment].Limits.Instant_buy[instant].Total.Amount)
            fmt.Println(payments[payment].Limits.Instant_buy[instant].Total.Currency)
            fmt.Println("payments[payment].Limits.Instant_buy[", instant, "].Remaining")
            fmt.Println(payments[payment].Limits.Instant_buy[instant].Remaining.Amount)
            fmt.Println(payments[payment].Limits.Instant_buy[instant].Remaining.Currency)
        }
        fmt.Println()
        for sell := range payments[payment].Limits.Sell {
            fmt.Println("payments[payment].Limits.Sell[", sell, "]")
            fmt.Println("payments[payment].Limits.Sell[", sell, "].Total")
            fmt.Println(payments[payment].Limits.Sell[sell].Period_in_days)
            fmt.Println(payments[payment].Limits.Sell[sell].Total.Amount)
            fmt.Println(payments[payment].Limits.Sell[sell].Total.Currency)
            fmt.Println("payments[payment].Limits.Sell[", sell, "].Remaining")
            fmt.Println(payments[payment].Limits.Sell[sell].Remaining.Amount)
            fmt.Println(payments[payment].Limits.Sell[sell].Remaining.Currency)
        }
        fmt.Println()
        for deposit := range payments[payment].Limits.Deposit {
            fmt.Println("payments[payment].Limits.Deposit[", deposit, "]")
            fmt.Println("payments[payment].Limits.Deposit[", deposit, "].Total")
            fmt.Println(payments[payment].Limits.Deposit[deposit].Period_in_days)
            fmt.Println(payments[payment].Limits.Deposit[deposit].Total.Amount)
            fmt.Println(payments[payment].Limits.Deposit[deposit].Total.Currency)
            fmt.Println("payments[payment].Limits.Deposit[", deposit, "].Remaining")
            fmt.Println(payments[payment].Limits.Deposit[deposit].Remaining.Amount)
            fmt.Println(payments[payment].Limits.Deposit[deposit].Remaining.Currency)
        }

        fmt.Println(payments[payment].Hold_business_days)
        fmt.Println(payments[payment].Hold_days)
        fmt.Println(payments[payment].Verificationmethod)
        fmt.Println(payments[payment].Cdvstatus)
        fmt.Println()
    }

    return payments
}

func cbp_get_all_transfers() []cbpPastTransfer {
    path := "/transfers"

    var transfers []cbpPastTransfer

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &transfers); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get all transfers")
    fmt.Println()
    for transfer := range transfers {
        fmt.Println("transfers[", transfer, "]")
        fmt.Println(transfers[transfer].Id)
        fmt.Println(transfers[transfer].Type)
        fmt.Println(transfers[transfer].Created_at)
        fmt.Println(transfers[transfer].Completed_at)
        fmt.Println(transfers[transfer].Canceled_at)
        fmt.Println(transfers[transfer].Processed_at)
        fmt.Println(transfers[transfer].Account_id)
        fmt.Println(transfers[transfer].User_id)
        fmt.Println(transfers[transfer].Amount)
        for k, v := range transfers[transfer].Details {
            fmt.Println(k, ":", v)
        }
        fmt.Println(transfers[transfer].User_nonce)
        fmt.Println()
    }

    return transfers
}

func cbp_get_transfer(transfer_id string) cbpPastTransfer {
    path := "/transfers/" + transfer_id

    var transfer cbpPastTransfer

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &transfer); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get a single transfer", transfer_id)
    fmt.Println()
    fmt.Println(transfer.Id)
    fmt.Println(transfer.Type)
    fmt.Println(transfer.Created_at)
    fmt.Println(transfer.Completed_at)
    fmt.Println(transfer.Canceled_at)
    fmt.Println(transfer.Processed_at)
    fmt.Println(transfer.Account_id)
    fmt.Println(transfer.User_id)
    fmt.Println(transfer.Amount)
    for k, v := range transfer.Details {
        fmt.Println(k, ":", v)
    }
    fmt.Println(transfer.User_nonce)
    fmt.Println()

    return transfer
}

func cbp_get_fee_estimate(currency string, crypto_address string) string {
    path := "/withdrawals/fee-estimate"

    var fee string

    response_status, response_body := cbp_rest_get_fee_estimate(path, currency, crypto_address)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &fee); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get fee estimate for crypto withdrawal")
    fmt.Println(fee)
    fmt.Println()

    return fee
}
