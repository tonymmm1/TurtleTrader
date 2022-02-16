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
    Account_id string `json:"account_id"`
    User_id string `json:"user_id"`
    Amount string `json:"amount"`
    Idem string `json:"idem"`
    Details map[string] interface {} `json:"details"`
    User_nonce string `json:"user_nonce"`
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
    Exchange_deposit_address bool `json:"exchange_deposit_address"`
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
        Min_withdrawal_amount float64 `json:"min_withdrawal_amount"`
        Max_withdrawal_amount float64 `json:"max_withdrawal_amount"`
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

type cbpTradingPair struct {
    Id string `json:"id"`
    Base_currency string `json:"base_currency"`
    Quote_currency string `json:"quote_currency"`
    Base_min_size string `json:"base_min_size"`
    Base_max_size string `json:"base_max_size"`
    Quote_increment string `json:"quote_increment"`
    Base_increment string `json:"base_increment"`
    Display_name string `json:"display_name"`
    Min_market_funds string `json:"min_market_funds"`
    Max_market_funds string `json:"max_market_funds"`
    Margin_enabled bool `json:"margin_enabled"`
    Post_only bool `json:"post_only"`
    Limit_only bool `json:limit_only"`
    Cancel_only bool `json:"cancel_only"`
    Status string `json:"status"`
    Status_message string `json:"status_message"`
    Trading_disabled bool `json:"trading_disabled"`
    Fx_stablecoin bool `json:"fx_stablecoin"`
    Max_slippage_percentage string `json:"max_slippage_percentage"`
    Auction_mode bool `json:"auction_mode"`
}

type cbpProductBook struct {
    Bids []interface{} `json:"bids"`
    Asks []interface{} `json:"asks"`
    Sequence float64 `json:"sequence"`
    Auction_mode bool `json:"auction_mode"`
    Auction struct {
        Open_price string `json:"open_price"`
        Open_size string `json:"open_size"`
        Best_bid_price string `json:"best_bid_price"`
        Best_bid_size string `json:"best_bid_size"`
        Best_ask_price string `json:"best_ask_price"`
        Best_ask_size string `json:"best_ask_size"`
        Auction_state string `json:"auction_state"`
        Can_open string `json:"can_open"`
        Time string `json:"time"`
    } `json:"auction"`
}

type cbpProductStats struct {
    Open string `json:"open"`
    High string `json:"high"`
    Low string `json:"low"`
    Last string `json:"last"`
    Volume string `json:"volume"`
    Volume_30day string `json:"volume_30day"`
}

type cbpProductTicker struct {
    Ask string `json:"ask"`
    Bid string `json:"bid"`
    Volume string `json:"volume"`
    Trade_id int32 `json:"trade_id"`
    Price string `json:"price"`
    Size string `json:"size"`
    Time string `json:"time"`
}

type cbpProductTrade struct {
    Trade_id int32 `json:"trade_id"`
    Side string `json:"side"`
    Size string `json:"size"`
    Price string `json:"price"`
    Time string `json:"time"`
}

/*  Accounts
*       Get all accounts for a profile      (GET)
*       Get a single account by id          (GET)
*       Get a single account's holds        (GET)
*       Get a single account's ledger       (GET)
*       Get a single account's transfers    (GET)
*/

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
    fmt.Println("Get all accounts for a profile")
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
    fmt.Println("Get a single account by id")
    fmt.Println()
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
    fmt.Println("Get a single account's holds")
    fmt.Println()
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

func cbp_get_single_account_ledger(account_id string) []cbpLedger { //List the holds of an account that belong to the same profile as the API key.
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
    fmt.Println("Get a single account's ledger")
    fmt.Println()
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
    fmt.Println("Get a single account's transfers")
    fmt.Println()
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

/*  Coinbase accounts
*       Get all Coinbase wallets    (GET)
*       Generate crypto address     (POST)
*/

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
    fmt.Println("Get all Coinbase wallets")
    fmt.Println()
    for wallet := range wallets {
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
        fmt.Println()
    }
    
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
    fmt.Println("Generate crypto address")
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
    fmt.Println(address.Exchange_deposit_address)
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

/*  Conversions
*       Convert currency    (POST)
*       Get a conversion    (GET)
*/

func cbp_convert_currency(profile_id string, from string, to string, amount string, nonce string) cbpConvert { //Converts funds from currency to currency
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
    fmt.Println()
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
    fmt.Println()
    fmt.Println(convert.Id)
    fmt.Println(convert.Amount)
    fmt.Println(convert.From_account_id)
    fmt.Println(convert.To_account_id)
    fmt.Println(convert.From)
    fmt.Println(convert.To)

    return convert
}

/*  Currencies
*       Get all known currencies    (GET)
*       Get a currency              (GET)
*/

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
        fmt.Println(currencies[currency].Details.Min_withdrawal_amount)
        fmt.Println(currencies[currency].Details.Max_withdrawal_amount)
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
    fmt.Println(currency.Details.Min_withdrawal_amount)
    fmt.Println(currency.Details.Max_withdrawal_amount)
    fmt.Println()

    return currency
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

func cbp_get_fee_estimate(currency string, crypto_address string) float64 {
    path := "/withdrawals/fee-estimate"
    
    var fee struct {
        Fee float64 `json:"fee"`
    }

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
    fmt.Println(fee.Fee)
    fmt.Println()

    return fee.Fee
}

/*  Fees
*       Get fees    (GET)
*/

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
    fmt.Println("Get fees")
    fmt.Println()
    fmt.Println(fees.Taker_fee_rate)
    fmt.Println(fees.Maker_fee_rate)
    fmt.Println(fees.Usd_volume)
    fmt.Println()

    return fees
}

/*  Orders
*       Get all fills       (GET)
*       Get all orders      (GET)
*       Cancel all orders   (DELETE)
*       Create a new order  (POST)
*       Get single order    (GET)
*       Cancel an order     (DELETE)
*/

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
    fmt.Println("Get all fills")
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
        fmt.Println()
    }

    return api_account_fills
}

/*  Coinbase price oracle
*       Get signed prices   (GET)
*/

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
    fmt.Print("Get signed prices")
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

/*  Products
*       Get all known trading pairs (GET)
*       Get single product          (GET)
*       Get product book            (GET)
*       Get product candles         (GET)
*       Get product stats           (GET)
*       Get product ticker          (GET)
*       Get product trades          (GET)
*/

func cbp_get_all_trading_pairs(query_type string) []cbpTradingPair {
    path := "/products"
    
    var trades []cbpTradingPair

    response_status, response_body := cbp_rest_get_all_trading_pairs(path, query_type)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &trades); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get all known trading pairs")
    fmt.Println()
    for trade := range trades {
        fmt.Println("trades[", trade, "]")
        fmt.Println(trades[trade].Id)
        fmt.Println(trades[trade].Base_currency)
        fmt.Println(trades[trade].Quote_currency)
        fmt.Println(trades[trade].Base_min_size)
        fmt.Println(trades[trade].Base_max_size)
        fmt.Println(trades[trade].Quote_increment)
        fmt.Println(trades[trade].Base_increment)
        fmt.Println(trades[trade].Display_name)
        fmt.Println(trades[trade].Min_market_funds)
        fmt.Println(trades[trade].Max_market_funds)
        fmt.Println(trades[trade].Margin_enabled)
        fmt.Println(trades[trade].Post_only)
        fmt.Println(trades[trade].Limit_only)
        fmt.Println(trades[trade].Cancel_only)
        fmt.Println(trades[trade].Status)
        fmt.Println(trades[trade].Status_message)
        fmt.Println(trades[trade].Trading_disabled)
        fmt.Println(trades[trade].Fx_stablecoin)
        fmt.Println(trades[trade].Max_slippage_percentage)
        fmt.Println(trades[trade].Auction_mode)
        fmt.Println()
    }

    return trades
}

func cbp_get_product(product_id string) cbpTradingPair {
    path := "/products/" + product_id

    var product cbpTradingPair

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &product); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get single product")
    fmt.Println()
    fmt.Println(product.Id)
    fmt.Println(product.Base_currency)
    fmt.Println(product.Quote_currency)
    fmt.Println(product.Base_min_size)
    fmt.Println(product.Base_max_size)
    fmt.Println(product.Quote_increment)
    fmt.Println(product.Base_increment)
    fmt.Println(product.Display_name)
    fmt.Println(product.Min_market_funds)
    fmt.Println(product.Max_market_funds)
    fmt.Println(product.Margin_enabled)
    fmt.Println(product.Post_only)
    fmt.Println(product.Limit_only)
    fmt.Println(product.Cancel_only)
    fmt.Println(product.Status)
    fmt.Println(product.Status_message)
    fmt.Println(product.Trading_disabled)
    fmt.Println(product.Fx_stablecoin)
    fmt.Println(product.Max_slippage_percentage)
    fmt.Println(product.Auction_mode)
    fmt.Println()

    return product
}

func cbp_get_product_book(product_id string, level int32) cbpProductBook {
    path := "/products/" + product_id + "/book"

    var book cbpProductBook

    response_status, response_body := cbp_rest_get_product_book(path, level)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &book); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get product book")
    fmt.Println()

    for k, v := range book.Bids {
        fmt.Println("book.Bids[", k, "]")
        fmt.Println(k, ":", v)
        fmt.Println()
    }
    for k, v := range book.Asks {
        fmt.Println("book.Asks[", k, "]")
        fmt.Println(k, ":", v)
        fmt.Println()
    }
    fmt.Println()
    fmt.Println(book.Sequence)
    fmt.Println(book.Auction_mode)
    fmt.Println(book.Auction.Open_price)
    fmt.Println(book.Auction.Best_bid_price)
    fmt.Println(book.Auction.Best_bid_size)
    fmt.Println(book.Auction.Best_ask_price)
    fmt.Println(book.Auction.Best_ask_size)
    fmt.Println(book.Auction.Auction_state)
    fmt.Println(book.Auction.Can_open)
    fmt.Println(book.Auction.Time)
    fmt.Println()

    return book
}

func cbp_get_product_candles(product_id string, granularity int32, start string, end string) []interface{} { //[]cbpProductCandle {
    path := "/products/" + product_id + "/candles"

    var candles []interface{}

    response_status, response_body := cbp_rest_get_product_candles(path, granularity, start, end)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &candles); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get product candles")
    fmt.Println()
    for k, v := range candles {
        fmt.Println(k, ":", v)
    }

    return candles
}

func cbp_get_product_stats(product_id string) cbpProductStats {
    path := "/products/" + product_id + "/stats"

    var stats cbpProductStats

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &stats); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get product stats")
    fmt.Println()
    fmt.Println(stats.Open)
    fmt.Println(stats.High)
    fmt.Println(stats.Low)
    fmt.Println(stats.Last)
    fmt.Println(stats.Volume)
    fmt.Println(stats.Volume_30day)
    fmt.Println()

    return stats
}

func cbp_get_product_ticker(product_id string) cbpProductTicker {
    path := "/products/" + product_id + "/ticker"

    var stats cbpProductTicker

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &stats); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get product ticker")
    fmt.Println()
    fmt.Println(stats.Ask)
    fmt.Println(stats.Bid)
    fmt.Println(stats.Volume)
    fmt.Println(stats.Trade_id)
    fmt.Println(stats.Price)
    fmt.Println(stats.Size)
    fmt.Println(stats.Time)
    fmt.Println()

    return stats
}

func cbp_get_product_trades(product_id string, limit int32) []cbpProductTrade {
    path := "/products/" + product_id + "/trades"

    var trades []cbpProductTrade

    response_status, response_body := cbp_rest_get_product_trades(path, limit)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &trades); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get product trades")
    fmt.Println()
    for trade := range trades {
        fmt.Println("trades[", trade, "]")
        fmt.Println(trades[trade].Trade_id)
        fmt.Println(trades[trade].Side)
        fmt.Println(trades[trade].Size)
        fmt.Println(trades[trade].Price)
        fmt.Println(trades[trade].Time)
        fmt.Println()
    }

    return trades
}

/*  Profiles
*       Get profiles                    (GET)
*       Create a profile                (POST)
*       Transfer funds between profiles (POST)
*       Get profile by id               (GET)
*       Rename a profile                (PUT)
*       Delete a profile                (PUT)
*/

func cbp_get_profiles() []cbpProfile{
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
    fmt.Println("Get profiles")
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

func cbp_create_profile(name string) cbpProfile {
    path := "/profiles"

    var profile cbpProfile

    response_status, response_body := cbp_rest_post_create_profile(path, name)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &profile); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }    

    //debug
    fmt.Println("Create a profile")
    fmt.Println()
    fmt.Println(profile.Id)
    fmt.Println(profile.User_id)
    fmt.Println(profile.Name)
    fmt.Println(profile.Active)
    fmt.Println(profile.Is_default)
    fmt.Println(profile.Has_margin)
    fmt.Println(profile.Created_at)
    fmt.Println()

    return profile
}

func cbp_transfer_funds_profiles(from string, to string, currency string, amount string) string {
    path := "/profiles/transfer"

    var profile string

    response_status, response_body := cbp_rest_post_transfer_funds_profiles(path, from, to, currency, amount)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    profile = string(response_body) //convert to string

    //debug
    fmt.Println("Transfer funds between profiles")
    fmt.Println()
    fmt.Println(profile)
    fmt.Println()

    return profile
}

func cbp_get_profile(profile_id string) cbpProfile {
    path := "/profiles/" + profile_id

    var profile cbpProfile

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &profile); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get profile by id")
    fmt.Println()
    fmt.Println(profile.Id)
    fmt.Println(profile.User_id)
    fmt.Println(profile.Name)
    fmt.Println(profile.Active)
    fmt.Println(profile.Is_default)
    fmt.Println(profile.Has_margin)
    fmt.Println(profile.Created_at)
    fmt.Println()

    return profile
}

/*  Reports
*       Get all reports (GET)
*       Create a report (POST)
*       Get a report    (GET)
*/

/*  Users
*       Get user exchange limits    (GET)    
*/
