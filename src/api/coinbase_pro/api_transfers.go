//Coinbase Pro API : Transfers
package coinbase_pro

import (
    "fmt"
    "os"
)

type Transfer struct { //Withdraw/deposit to/from Coinbase/payment
    Id string `json:"id"`
    Amount string `json:"amount"`
    Currency string `json:"currency"`
    Payout_at string `json:"payout_at"`
    Fee string `json:"fee"`
    Subtotal string `json:"subtotal"`
}

type Payment struct { //Get all payment methods
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

func transfer_coinbase_account(profile_id string, amount string, account_id string, currency string) []byte {
    path := "/deposits/coinbase-account"

    status, response := rest_post_coinbase(path, profile_id, amount, account_id, currency)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func transfer_payment_account(profile_id string, amount string, account_id string, currency string) []byte {
    path := "/deposits/payment-method"

    status, response := rest_post_payment(path, profile_id, amount, account_id, currency)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_all_payments() []byte {
    path := "/payment-methods"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_all_transfers() []byte {
    path := "/transfers"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_transfer(transfer_id string) []byte {
    path := "/transfers/" + transfer_id

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_fee_estimate(currency string, crypto_address string) []byte {
    path := "/withdrawals/fee-estimate"

    status, response := rest_get_fee_estimate(path, currency, crypto_address)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
