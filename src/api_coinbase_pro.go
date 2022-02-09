//Coinbase Pro API handler functions
package main

import (
        "encoding/json"
        "fmt"
        "os"
)

const ( //Common return codes (https://docs.cloud.coinbase.com/exchange/docs/requests)
    CBP_STATUS_CODE_SUCCESS int = 200           //Success
    CBP_STATUS_CODE_BAD_REQUEST = 400           //Bad Request -- Invalid request format
    CBP_STATUS_CODE_UNAUTHORIZED = 401          //Unauthorized -- Invalid API Key
    CBP_STATUS_CODE_FORBIDDEN = 403             //Forbidden -- You do not have access to the requested resource
    CBP_STATUS_CODE_NOT_FOUND = 404             //Not Found
    CBP_STATUS_CODE_INTERNAL_SERVER_ERROR = 500 //Internal Server Error -- We had a problem with our server
)

type cbpConfig struct { //Coinbase Pro configuration
    Host string
    Key string
    Password string
    Secret string
}

type cbpAccount struct { //Get all accounts for a profile/Get a single account by id
    Id string `json:"id"`
    Currency string `json:"currency"`
    Balance string `json:"balance"`
    Hold string `json:"hold"`
    Available string `json:"available"`
    Profile_id string `json:"profile_id"`
    Trading_enabled bool `json:"trading_enabled"`
}

type cbpLedger struct { //struct to store API ledger
    Id string `json:"id"`
    Amount string `json:"amount"`
    Balance string `json:"balance"`
    Created_at string `json:"created_at"`
    Type string `json:"type"`
    Details map[string] interface {} `json:"details"`
}

type cbpHold struct { //struct to store API hold
    Id string `json:"id"`
    Created_at string `json:"created_at"`
    Amount string `json:"amount"`
    Ref string `json:"ref"`
    Type string `json:"type"`
}

type cbpPastTransfer struct { //struct to store API past transfer
    Id string `json:"id"`
    Type string `json:"type"`
    Created_at string `json:"created_at"`
    Completed_at string `json:"completed_at"`
    Canceled_at string `json:"canceled_at"`
    Processed_at string `json:"processed_at"`
    User_nonce string `json:"user_nonce"`
    Amount string `json:"amount"`
    Details map[string] interface {} `json:"details"`
    Idem string `json:"idem"`
}

type cbpWallet struct { //struct to store API wallet
    Id string `json:"id"`
    Name string `json:"name"`
    Balance string `json:"balance"`
    Currency string `json:"currency"`
    Type string `json:"type"`
    Primary bool `json:"primary"`
    Active bool `json:"active"`
    Available_on_consumer bool `json:"available_on_consumer"`
    //Ready bool `json:"ready"`
    Wire_deposit_information struct {
        Account_number string `json:"account_number"`
        Routing_number string `json:"routing_number"`
        Bank_name string `json:"bank_name"`
        Bank_address string `json:"bank_address"`
        Bank_country struct {
            Code string `json:"code"`
            Name string `json:"name"`
        } `json:"bank_country"`
        Account_name string `json:"account_name"`
        Account_address string `json:"account_address"`
        Reference string `json:"reference"`
    } `json:"wire_deposit_information"`
    Swift_deposit_information struct {
        Account_number string `json:"account_number"`
        Routing_number string `json:"routing_number"`
        Bank_name string `json:"bank_name"`
        Bank_address string `json:"bank_address"`
        Bank_country struct {
            Code string `json:"code"`
            Name string `json:"name"`
        } `json:"bank_country"`
        Account_name string `json:"account_name"`
        Account_address string `json:"account_address"`
        Reference string `json:"reference"`
    } `json:"swift_deposit_information"`
    Sepa_deposit_information struct {
        Iban string `json:"iban"`
        Swift string `json:"swift"`
        Bank_name string `json:"bank_name"`
        Bank_address string `json:"bank_address"`
        Bank_country_name string `json:"bank_country_name"`
        Account_name string `json:"account_name"`
        Account_address string `json:"account_address"`
        Reference string `json:"reference"`
    } `json:"sepa_deposit_information"`
    Uk_deposit_information struct {
        Sort_code string `json:"sort_code"`
        Account_number string `json:"account_number"`
        Bank_name string `json:"bank_name"`
        Account_name string `json:"account_name"`
        Reference string `json:"reference"`
    } `json:"uk_deposit_information"`
    //Destination_tag_name string `json:"destination_tag_name"`
    //Destination_tag_regex string `json:"destination_tag_regex"`
    Hold_balance string `json:"hold_balance"`
    Hold_currency string `json:"hold_currency"`
}

type cbpCryptoAddress struct { //struct to store API generated crypto address
    Id string `json:"id"`
    Address string `json:"address"`
    Address_info struct {
        Address string `json:"address"`
        Destination_tag string `json:"destination_tag"`
    } `json:"address_info"`
    Name string `json:"name"`
    Created_at string `json:"created_at"`
    Updated_at string `json:"updated_at"`
    Network string `json:"network"`
    Uri_scheme string `json:"uri_scheme"`
    Resource string `json:"resource"`
    Resource_path string `json:"resource_path"`
    Warnings []struct {
        Title string `json:"title"`
        Details string `json:"details"`
        Image_url string `json:"image_url"`
    } `json:"warnings"`
    //Legacy_address string `json:"legacy_address"`
    //Destination_tag string `json:"destination_tag"`
    //Deposit_uri string `json:"deposity_uri"`
    //Callback_url string `json:"callback_url"`
}

type cbpCurrency struct { //Get all known currency/Get a currency
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
        Min_withdrawl_amount float64 `json:"min_withdrawl_amount"`
        Max_withdrawl_amount float64 `json:"max_withdrawl_amount"`
    } `json:"details"`
}

type cbpConvert struct { //Convert Currency/Get a conversion
    Id string `json:"id"`
    Amount string `json:"amount"`
    From_account_id string `json:"from_account_id"`
    To_account_id string `json:"to_account_id"`
    From string `json:"from"`
    To string `json:"to"`
}

type cbpFee struct { //Get fees
    Taker_fee_rate string `json:"taker_fee_rate"`
    Maker_fee_rate string `json:"maker_fee_rate"`
    Usd_volume string `json:"usd_volume"`
}

type cbpTransfer struct { //Withdraw/deposit to/from Coinbase/payment
    Id string `json:"id"`
    Amount string `json:"amount"`
    Currency string `json:"currency"`
    Payout_at string `json:"payout_at"`
    Fee string `json:"fee"`
    Subtotal string `json:"subtotal"`
}

type cbpFill struct { //Get all fills
    Trade_id int32 `json:"id"`
    Product_id string `json:"product_id"`
    Order_id string `json:"order_id"`
    User_id string `json:"user_id"`
    Profile_id string `json:"profile_id"`
    Liquidity string `json:"liquidity"`
    Price string `json:"price"`
    Size string `json:"size"`
    Fee string `json:"fee"`
    Created_at string `json:"created_at"`
    Side string `json:"side"`
    Settled bool `json:"settled"`
    Usd_volume string `json:"usd_volume"`
}

type cbpProfile struct { //Get profiles/Create a profile/Get profile by id/Get profile by id
    Id string `json:"id"`
    User_id string `json:"user_id"`
    Name string `json:"name"`
    Active bool `json:"active"`
    Is_default bool `json:"is_default"`
    Has_margin bool `json:"has_margin"`
    Created_at string `json:"created_at"`
}

type cbpPrice struct { //Get signed prices
    Timestamp string `json:"timestamp"`
    Messages []string `json:"messages"`
    Signatures []string `json:"signatures"`
    Prices map[string] interface {} `json:"prices"`
}

func cbp_get_all_accounts() []cbpAccount { //Get a list of trading accounts from the profile of the API key.
    path := "/accounts"

    var accounts []cbpAccount

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &accounts); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("accounts:")
    fmt.Println()
    for account := range accounts {
        fmt.Println("accounts[", account, "]")
        fmt.Println(accounts[account].Id)
        fmt.Println(accounts[account].Currency)
        fmt.Println(accounts[account].Balance)
        fmt.Println(accounts[account].Hold)
        fmt.Println(accounts[account].Available)
        fmt.Println(accounts[account].Profile_id)
        fmt.Println(accounts[account].Trading_enabled)
        fmt.Println()
    }

    return accounts
}

func cbp_get_single_account(account_id string) cbpAccount { //Information for a single account.
    path := "/accounts/" + account_id

    var account cbpAccount //store single cbpAccount

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &account); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("account:")
    fmt.Println(account.Id)
    fmt.Println(account.Currency)
    fmt.Println(account.Balance)
    fmt.Println(account.Hold)
    fmt.Println(account.Available)
    fmt.Println(account.Profile_id)
    fmt.Println(account.Trading_enabled)
    fmt.Println()

    return account
}

func cbp_get_single_account_holds(account_id string) []cbpHold { //List the holds of an account that belong to the same profile as the API key.
    path := "/accounts/" + account_id + "/holds" //?limit=100" //implement limit logic later

    var holds []cbpHold

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }
    if response_body == nil {
        return nil
    }

    if err := json.Unmarshal(response_body, &holds); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("holds:")
    for hold := range holds {
        fmt.Println("holds[", hold,"]")
        fmt.Println(holds[hold].Id)
        fmt.Println(holds[hold].Created_at)
        fmt.Println(holds[hold].Amount)
        fmt.Println(holds[hold].Ref)
        fmt.Println(holds[hold].Type)
        fmt.Println()
    }

    return holds
}

func cbp_get_single_account_ledgers(account_id string) []cbpLedger { //List the holds of an account that belong to the same profile as the API key.
    path := "/accounts/" + account_id + "/ledger" //?limit=100" //implement limit logic later

    var ledgers []cbpLedger

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &ledgers); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("account_id:", account_id)
    for ledger := range ledgers {
        fmt.Println(ledgers[ledger].Id)
        fmt.Println(ledgers[ledger].Amount)
        fmt.Println(ledgers[ledger].Balance)
        fmt.Println(ledgers[ledger].Created_at)
        fmt.Println(ledgers[ledger].Type)
        for k, v := range ledgers[ledger].Details {
            fmt.Println(k, ":", v)
        }
        fmt.Println()
    }

    return ledgers
}

func cbp_get_single_account_transfers(account_id string) []cbpPastTransfer { //Lists past withdrawals and deposits for an account.
    path := "/accounts/" + account_id + "/transfers" //?limit=100" //implement limit logic later

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
    fmt.Println("account_id:", account_id)
    for transfer := range transfers {
        fmt.Println(transfers[transfer].Id)
        fmt.Println(transfers[transfer].Type)
        fmt.Println(transfers[transfer].Created_at)
        fmt.Println(transfers[transfer].Completed_at)
        fmt.Println(transfers[transfer].Canceled_at)
        fmt.Println(transfers[transfer].Processed_at)
        fmt.Println(transfers[transfer].User_nonce)
        fmt.Println(transfers[transfer].Amount)
        for k, v := range transfers[transfer].Details {
            fmt.Println(k, ":", v)
        }
        fmt.Println()
    }

    return transfers 
}

func cbp_get_all_wallets() []cbpWallet { //Gets all the user's available Coinbase wallets
    path := "/coinbase-accounts"

    var wallets []cbpWallet

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &wallets); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("api_account_wallets: ", wallets)
    fmt.Println()
    for wallet := range wallets {
        fmt.Println()
        fmt.Println("wallets[", wallet, "]")
        fmt.Println(wallets[wallet].Id)
        fmt.Println(wallets[wallet].Name)
        fmt.Println(wallets[wallet].Balance)
        fmt.Println(wallets[wallet].Currency)
        fmt.Println(wallets[wallet].Type)
        fmt.Println(wallets[wallet].Primary)
        fmt.Println(wallets[wallet].Active)
        fmt.Println(wallets[wallet].Available_on_consumer)
        /*if wallets[wallet].Ready == (true || false) {
            fmt.Println(wallets[wallet].Ready)
        }*/
        if wallets[wallet].Wire_deposit_information.Account_name != "" {
            fmt.Println(wallets[wallet].Wire_deposit_information.Account_number)
            fmt.Println(wallets[wallet].Wire_deposit_information.Routing_number)
            fmt.Println(wallets[wallet].Wire_deposit_information.Bank_name)
            fmt.Println(wallets[wallet].Wire_deposit_information.Bank_address)
            fmt.Println(wallets[wallet].Wire_deposit_information.Bank_country.Code)
            fmt.Println(wallets[wallet].Wire_deposit_information.Bank_country.Name)
            fmt.Println(wallets[wallet].Wire_deposit_information.Account_name)
            fmt.Println(wallets[wallet].Wire_deposit_information.Account_address)
            fmt.Println(wallets[wallet].Wire_deposit_information.Reference)
        }
        if wallets[wallet].Swift_deposit_information.Account_name != "" {
            fmt.Println(wallets[wallet].Swift_deposit_information.Account_number)
            fmt.Println(wallets[wallet].Swift_deposit_information.Routing_number)
            fmt.Println(wallets[wallet].Swift_deposit_information.Bank_name)
            fmt.Println(wallets[wallet].Swift_deposit_information.Bank_address)
            fmt.Println(wallets[wallet].Swift_deposit_information.Bank_country.Code)
            fmt.Println(wallets[wallet].Swift_deposit_information.Bank_country.Name)
            fmt.Println(wallets[wallet].Swift_deposit_information.Account_name)
            fmt.Println(wallets[wallet].Swift_deposit_information.Account_address)
            fmt.Println(wallets[wallet].Swift_deposit_information.Reference)
        }
        if wallets[wallet].Sepa_deposit_information.Account_name != ""{
            fmt.Println(wallets[wallet].Sepa_deposit_information.Iban)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Swift)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Bank_name)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Bank_address)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Bank_country_name)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Account_name)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Account_address)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Reference)
        }
        if wallets[wallet].Uk_deposit_information.Account_name != "" {
            fmt.Println(wallets[wallet].Uk_deposit_information.Sort_code)
            fmt.Println(wallets[wallet].Uk_deposit_information.Account_number)
            fmt.Println(wallets[wallet].Uk_deposit_information.Bank_name)
            fmt.Println(wallets[wallet].Uk_deposit_information.Account_name)
            fmt.Println(wallets[wallet].Uk_deposit_information.Reference)
        }
        /*if wallets[wallet].Destination_tag_name != "" {
            fmt.Println(wallets[wallet].Destination_tag_name)
        }
        if wallets[wallet].Destination_tag_regex != "" {
            fmt.Println(wallets[wallet].Destination_tag_regex)
        }*/
        fmt.Println(wallets[wallet].Hold_balance)
        fmt.Println(wallets[wallet].Hold_currency)
    }
    fmt.Println()
    
    return wallets
}

func cbp_generate_crypto_address(account_id string) cbpCryptoAddress { //Generates a one-time crypto address for depositing crypto.
    path := "/coinbase-accounts/" + account_id + "/addresses"

    var address cbpCryptoAddress

    response_status, response_body := cbp_rest_post_address(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &address); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("address:", account_id)
    fmt.Println()
    fmt.Println(address.Id)
    fmt.Println(address.Address)
    if address.Address_info.Address != "" {
        fmt.Println(address.Address_info.Address)
        fmt.Println(address.Address_info.Destination_tag)
    }
    fmt.Println(address.Name)
    fmt.Println(address.Created_at)
    fmt.Println(address.Updated_at)
    if address.Network != "" {
        fmt.Println(address.Network)
    }
    if address.Uri_scheme != "" {
        fmt.Println(address.Uri_scheme)
    }
    fmt.Println(address.Resource)
    fmt.Println(address.Resource_path)
    if address.Warnings != nil {
        for warning := range address.Warnings {
            fmt.Println()
            fmt.Println("address.Warnings[", warning, "]")
            fmt.Println(address.Warnings[warning].Title)
            fmt.Println(address.Warnings[warning].Details)
            fmt.Println(address.Warnings[warning].Image_url)
        }
    }
    /*if address.Legacy_address != "" {
        fmt.Println(address.Legacy_address)
    }
    if address.Destination_tag != "" {
        fmt.Println(address.Destination_tag)
    }
    if address.Deposit_uri != "" {
        fmt.Println(address.Deposit_uri)
    }
    if address.Callback_url != "" {
        fmt.Println(address.Callback_url)
    }*/
    fmt.Println()

    return address
}

func cbp_get_all_currencies() []cbpCurrency {
    path := "/currencies"

    var currencies []cbpCurrency

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
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
        fmt.Println(currencies[currency].Details.Min_withdrawl_amount)
        fmt.Println(currencies[currency].Details.Max_withdrawl_amount)
        fmt.Println()
    }

    return currencies
}

func cbp_get_currency(currency_id string) cbpCurrency {
    path := "/currencies/" + currency_id

    var currency cbpCurrency

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
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
    fmt.Println(currency.Details.Min_withdrawl_amount)
    fmt.Println(currency.Details.Max_withdrawl_amount)
    fmt.Println()

    return currency
}

func cbp_convert_currency(profile_id string, from string, to string, amount string, nonce string) cbpConvert { //Converts funs from currency to currency
    path := "/conversions"

    var convert cbpConvert

    response_status, response_body := cbp_rest_post_convert(path, profile_id, from, to, amount, nonce)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &convert); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Convert currency")
    fmt.Println(convert.Id)
    fmt.Println(convert.Amount)
    fmt.Println(convert.From_account_id)
    fmt.Println(convert.To_account_id)
    fmt.Println(convert.From)
    fmt.Println(convert.To)

    return convert
}

func cbp_get_conversion(conversion_id string, profile_id string) cbpConvert{
    path := "/conversion/" + conversion_id

    var convert cbpConvert

    response_status, response_body := cbp_rest_get_convert(path, profile_id)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &convert); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get a conversion")
    fmt.Println(convert.Id)
    fmt.Println(convert.Amount)
    fmt.Println(convert.From_account_id)
    fmt.Println(convert.To_account_id)
    fmt.Println(convert.From)
    fmt.Println(convert.To)

    return convert
}

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
    fmt.Println(deposit.Id)
    fmt.Println(deposit.Amount)
    fmt.Println(deposit.Currency)
    fmt.Println(deposit.Payout_at)
    fmt.Println(deposit.Fee)
    fmt.Println(deposit.Subtotal)

    return deposit
}

func cbp_get_fees() cbpFee {
    path := "/fees"

    var fees cbpFee

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &fees); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("fees:")
    fmt.Println(fees.Taker_fee_rate)
    fmt.Println(fees.Maker_fee_rate)
    fmt.Println(fees.Usd_volume)
    fmt.Println()

    return fees
}

func cbp_get_all_fills(order_id string, product_id string, profile_id string, limit int64, before int64, after int64) []cbpFill {
    path := "/fills"

    var api_account_fills []cbpFill

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_account_fills); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("api_account_fills:")
    fmt.Println()
    for fill := range api_account_fills {
        fmt.Println("api_account_fills[", fill, "]")
        fmt.Println(api_account_fills[fill].Trade_id)
        fmt.Println(api_account_fills[fill].Product_id)
        fmt.Println(api_account_fills[fill].Order_id)
        fmt.Println(api_account_fills[fill].User_id)
        fmt.Println(api_account_fills[fill].Profile_id)
        fmt.Println(api_account_fills[fill].Liquidity)
        fmt.Println(api_account_fills[fill].Price)
        fmt.Println(api_account_fills[fill].Size)
        fmt.Println(api_account_fills[fill].Fee)
    }
    fmt.Println()

    return api_account_fills
}

func cbp_get_profiles(active bool) []cbpProfile{
    path := "/profiles"

    var profiles []cbpProfile

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &profiles); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("api_account_profiles:")
    fmt.Println()
    for profile := range profiles {
        fmt.Println("profiles[", profile, "]")
        fmt.Println(profiles[profile].Id)
        fmt.Println(profiles[profile].User_id)
        fmt.Println(profiles[profile].Name)
        fmt.Println(profiles[profile].Active)
        fmt.Println(profiles[profile].Is_default)
        fmt.Println(profiles[profile].Has_margin)
        fmt.Println(profiles[profile].Created_at)
    }
    fmt.Println()

    return profiles
}

func cbp_get_signed_prices() cbpPrice {
    path := "/oracle"

    var prices cbpPrice

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &prices); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("api_account_prices:")
    fmt.Println()
    fmt.Println(prices.Timestamp)
    for message := range prices.Messages {
        fmt.Println("prices.Messages[", message, "]")
        fmt.Println(prices.Messages[message])
        fmt.Println()
    }
    fmt.Println()
    for signature := range prices.Signatures {
        fmt.Println("price.Signatures[", signature, "]")
        fmt.Println(prices.Signatures[signature])
        fmt.Println()
    }
    fmt.Println()
    for k, v := range prices.Prices {
        fmt.Println(k, ":", v)
    }
    fmt.Println()

    return prices
}
