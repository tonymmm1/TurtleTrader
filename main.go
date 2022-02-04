package main

import (
        "crypto/hmac"
        "crypto/sha256"
        "encoding/base64"
        "encoding/json"
        "fmt"
        "os"
        "strconv"
        "time"

        "github.com/BurntSushi/toml"
        "github.com/go-resty/resty/v2"
)

const ( //Common return codes (https://docs.cloud.coinbase.com/exchange/docs/requests)
    STATUS_CODE_SUCCESS int = 200           //Success
    STATUS_CODE_BAD_REQUEST = 400           //Bad Request -- Invalid request format
    STATUS_CODE_UNAUTHORIZED = 401          //Unauthorized -- Invalid API Key
    STATUS_CODE_FORBIDDEN = 403             //Forbidden -- You do not have access to the requested resource
    STATUS_CODE_NOT_FOUND = 404             //Not Found
    STATUS_CODE_INTERNAL_SERVER_ERROR = 500 //Internal Server Error -- We had a problem with our server
)

const ( //POST request types
    POST_REQUEST_GENERATE_ADDRESS = 1
    POST_REQUEST_CONVERT_CURRENCY = 2
    POST_REQUEST_CREATE_ORDER = 3
    POST_REQUEST_DEPOSIT_FROM_COINBASE = 4
    POST_REQUEST_DEPOSIT_FROM_PAYMENT = 5
    POST_REQUEST_WITHDRAW_TO_COINBASE = 6
    POST_REQUEST_WITHDRAW_TO_CRYPTO = 7
    POST_REQUEST_WITHDRAW_TO_PAYMENT = 8
    POST_REQUEST_CREATE_PROFILE = 9
    POST_REQUEST_TRANSFER_FUNDS_TO_PROFILE = 10
    POST_REQUEST_CREATE_REPORT = 11
)

type apiConfig struct { //configuration toml file struct
    Host string
    Key string
    Password string
    Secret string
}

type apiAccount struct { //struct to store API account
    Id string `json:"id"`
    Currency string `json:"currency"`
    Balance string `json:"balance"`
    Hold string `json:"hold"`
    Available string `json:"available"`
    Profile_id string `json:"profile_id"`
    Trading_enabled bool `json:"trading_enabled"`
}

type apiLedger struct { //struct to store API ledger
    Id string `json:"id"`
    Amount string `json:"amount"`
    Balance string `json:"balance"`
    Created_at string `json:"created_at"`
    Type string `json:"type"`
    Details struct {
        Order_id string `json:"order_id"`
        Product_id string `json:"product_id"`
        Trade_id string `json:"trade_id"`
    } `json:"details"`
}

type apiHold struct { //struct to store API hold
    Id string `json:"id"`
    Created_at string `json:"created_at"`
    Amount string `json:"amount"`
    Ref string `json:"ref"`
    Type string `json:"type"`
}

type apiPastTransfer struct { //struct to store API past transfer
    Id string `json:"id"`
    Type string `json:"type"`
    Created_at string `json:"created_at"`
    Completed_at string `json:"completed_at"`
    Canceled_at string `json:"canceled_at"`
    Processed_at string `json:"processed_at"`
    User_nonce string `json:"user_nonce"`
    Amount string `json:"amount"`
    Details struct {
        Coinbase_payout_at string `json:"coinbase_payout_at"`
        Coinbase_account_id string `json:"coinbase_account_id"`
        Coinbase_deposit_id string `json:"coinbase_deposit_id"`
        Coinbase_payment_method_id string `json:"coinbase_payment_method_id"`
        Coinbase_payment_method_type string `json:"coinbase_payment_method_type"`
    } `json:"details"`
    Idem string `json:"idem"`
}

type apiWallet struct { //struct to store API wallet
    Id string `json:"id"`
    Name string `json:"name"`
    Balance string `json:"balance"`
    Currency string `json:"currency"`
    Type string `json:"type"`
    Primary bool `json:"primary"`
    Active bool `json:"active"`
    Available_on_consumer bool `json:"available_on_consumer"`
    Ready bool `json:"ready"`
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
    Destination_tag_name string `json:"destination_tag_name"`
    Destination_tag_regex string `json:"destination_tag_regex"`
    Hold_balance string `json:"hold_balance"`
    Hold_currency string `json:"hold_currency"`
}

type apiCryptoAddress struct { //struct to store API generated crypto address
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
    Legacy_address string `json:"legacy_address"`
    Destination_tag string `json:"destination_tag"`
    Deposit_uri string `json:"deposity_uri"`
    Callback_url string `json:"callback_url"`
}

type reqConvert struct {
    Request_path string
    Profile_id string
    From string
    To string
    Amount string
    Nonce string
}

type apiConvert struct { //struct to store API conversion
    Id string `json:"id"`
    Amount string `json:"amount"`
    From_account_id string `json:"from_account_id"`
    To_account_id string `json:"to_account_id"`
    From string `json:"from"`
    To string `json:"to"`
}

type reqTransfer struct {
    Request_path string
    Profile_id string
    Amount string
    Payment_method_id string
    Coinbase_account_id string
    Currency string
    Crypto_address string   //crypto address
    Destination_tag string  //crypto address
    No_destination_tag bool //crypto address
    Two_factor_code string  //crypto address
    Nonce int32 //unique value //crypto address
    Fee string  //set with post_get_fee_estimate //crypto address
}

type apiFee struct {
    Taker_fee_rate string `json:"taker_fee_rate"`
    Maker_fee_rate string `json:"maker_fee_rate"`
    Usd_volume string `json:"usd_volume"`
}

type apiTransfer struct {
    Id string `json:"id"`
    Amount string `json:"amount"`
    Currency string `json:"currency"`
    Payout_at string `json:"payout_at"`
    Fee string `json:"fee"`
    Subtotal string `json:"subtotal"`
}

type reqFill struct {
    Request_path string
    Order_id string
    Product_id string
    Profile_id string
    Limit int64
    Before int64
    After int64
}

type apiFill struct {
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

type apiProfile struct {
    Id string `json:"id"`
    User_id string `json:"user_id"`
    Name string `json:"name"`
    Active bool `json:"active"`
    Is_default bool `json:"is_default"`
    Has_margin bool `json:"has_margin"`
    Created_at string `json:"created_at"`
}

type apiPrice struct {
    Timestamp string `json:"timestamp"`
    Messages []string `json:"messages"`
    Signatures []string `json:"signatures"`
    Prices map[string] interface {} `json:"prices"`
}

func gen_api_message(api_key_secret string, time_current string, request_method string, request_path string) string { //generate hashed message for REST requests
    message := time_current + request_method + request_path //construct prehase message

    decoded_secret, err := base64.StdEncoding.DecodeString(api_key_secret) //decode base64 encoded api secret
    if err != nil {
        fmt.Println("ERROR decoding api key secret")
        os.Exit(1)
    }

    hash := hmac.New(sha256.New, []byte(decoded_secret)) //generate new SHA256 hmac based on decoded api secret
    hash.Write([]byte(message)) //hash message using hmac

    return base64.StdEncoding.EncodeToString(hash.Sum(nil)) //return hashed message
}

func rest_get(api_struct apiConfig, request_path string) (int, []byte) { //handles GET requests
    request_method := "GET"

    time_current := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message_hashed := gen_api_message(api_struct.Secret, time_current, request_method, request_path) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : api_struct.Key,
            "CB-ACCESS-SIGN" : message_hashed,
            "CB-ACCESS-TIMESTAMP" : time_current,
            "CB-ACCESS-PASSPHRASE" : api_struct.Password,
            "Content-Type" : "application/json"}).
        SetAuthToken(api_struct.Key).
        Get(api_struct.Host + request_path)

    // debug
    fmt.Println("Response Info:")
    fmt.Println("  Error      :", err)
    fmt.Println("  Status Code:", resp.StatusCode())
    fmt.Println("  Status     :", resp.Status())
    fmt.Println("  Proto      :", resp.Proto())
    fmt.Println("  Time       :", resp.Time())
    fmt.Println("  Received At:", resp.ReceivedAt())
    fmt.Println("  Body       :\n", resp)
    fmt.Println()

    return resp.StatusCode(), resp.Body()
}

func rest_get_all_fills(api_struct apiConfig, request_struct reqFill) (int, []byte) { //handles GET requests
    request_method := "GET"

    time_current := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message_hashed := gen_api_message(api_struct.Secret, time_current, request_method, request_struct.Request_path) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : api_struct.Key,
            "CB-ACCESS-SIGN" : message_hashed,
            "CB-ACCESS-TIMESTAMP" : time_current,
            "CB-ACCESS-PASSPHRASE" : api_struct.Password,
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "order_id" : request_struct.Order_id,
            "product_id" : request_struct.Product_id,
            "profile_id" : request_struct.Profile_id,
            "limit" : strconv.FormatInt(request_struct.Limit, 10),
            "before" : strconv.FormatInt(request_struct.Before, 10),
            "after" : strconv.FormatInt(request_struct.After, 10)}).
        SetAuthToken(api_struct.Key).
        Get(api_struct.Host + request_struct.Request_path)

    // debug
    fmt.Println("Response Info:")
    fmt.Println("  Error      :", err)
    fmt.Println("  Status Code:", resp.StatusCode())
    fmt.Println("  Status     :", resp.Status())
    fmt.Println("  Proto      :", resp.Proto())
    fmt.Println("  Time       :", resp.Time())
    fmt.Println("  Received At:", resp.ReceivedAt())
    fmt.Println("  Body       :\n", resp)
    fmt.Println()

    return resp.StatusCode(), resp.Body()
}

func rest_get_profiles(api_struct apiConfig, request_path string, active bool) (int, []byte) {
    request_method := "GET"

    time_current := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message_hashed := gen_api_message(api_struct.Secret, time_current, request_method, request_path) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : api_struct.Key,
            "CB-ACCESS-SIGN" : message_hashed,
            "CB-ACCESS-TIMESTAMP" : time_current,
            "CB-ACCESS-PASSPHRASE" : api_struct.Password,
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "active" : strconv.FormatBool(active)}).
        SetAuthToken(api_struct.Key).
        Get(api_struct.Host + request_path)

    // debug
    fmt.Println("Response Info:")
    fmt.Println("  Error      :", err)
    fmt.Println("  Status Code:", resp.StatusCode())
    fmt.Println("  Status     :", resp.Status())
    fmt.Println("  Proto      :", resp.Proto())
    fmt.Println("  Time       :", resp.Time())
    fmt.Println("  Received At:", resp.ReceivedAt())
    fmt.Println("  Body       :\n", resp)
    fmt.Println()

    return resp.StatusCode(), resp.Body()
}

func rest_post_generate_address(api_struct apiConfig, request_path string) (int, []byte) { //POST_REQUEST_GENERATE_ADDRESS
    request_method := "POST"

    time_current := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message_hashed := gen_api_message(api_struct.Secret, time_current, request_method, request_path) //create hashed message to send

    client := resty.New() //create REST session

    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : api_struct.Key,
            "CB-ACCESS-SIGN" : message_hashed,
            "CB-ACCESS-TIMESTAMP" : time_current,
            "CB-ACCESS-PASSPHRASE" : api_struct.Password,
            "Content-Type" : "application/json"}).
        SetAuthToken(api_struct.Key).
        Post(api_struct.Host + request_path)

    // debug
    fmt.Println("Response Info:")
    fmt.Println("  Error      :", err)
    fmt.Println("  Status Code:", resp.StatusCode())
    fmt.Println("  Status     :", resp.Status())
    fmt.Println("  Proto      :", resp.Proto())
    fmt.Println("  Time       :", resp.Time())
    fmt.Println("  Received At:", resp.ReceivedAt())
    fmt.Println("  Body       :\n", resp)
    fmt.Println()
    
    return resp.StatusCode(), resp.Body()
}

func rest_post_convert_currency(api_struct apiConfig, request_struct reqConvert) (int, []byte) { //POST_REQUEST_CONVERT_CURRENCY
    request_method := "POST"

    time_current := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message_hashed := gen_api_message(api_struct.Secret, time_current, request_method, request_struct.Request_path) //create hashed message to send

    client := resty.New() //create REST session

    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : api_struct.Key,
            "CB-ACCESS-SIGN" : message_hashed,
            "CB-ACCESS-TIMESTAMP" : time_current,
            "CB-ACCESS-PASSPHRASE" : api_struct.Password,
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "profile_id" : request_struct.Profile_id,
            "from" : request_struct.From,
            "to" : request_struct.To,
            "amount" : request_struct.Amount,
            "nonce" : request_struct.Nonce}).
        SetAuthToken(api_struct.Key).
        Post(api_struct.Host + request_struct.Request_path)

    // debug
    fmt.Println("Response Info:")
    fmt.Println("  Error      :", err)
    fmt.Println("  Status Code:", resp.StatusCode())
    fmt.Println("  Status     :", resp.Status())
    fmt.Println("  Proto      :", resp.Proto())
    fmt.Println("  Time       :", resp.Time())
    fmt.Println("  Received At:", resp.ReceivedAt())
    fmt.Println("  Body       :\n", resp)
    fmt.Println()

    return resp.StatusCode(), resp.Body()
}
func rest_post_transfer_payment(api_struct apiConfig, request_struct reqTransfer) (int, []byte) { //POST_REQUEST_(WITHDRAW/DEPOSIT)_PAYMENT
    request_method := "POST"

    time_current := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message_hashed := gen_api_message(api_struct.Secret, time_current, request_method, request_struct.Request_path) //create hashed message to send

    client := resty.New() //create REST session

    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : api_struct.Key,
            "CB-ACCESS-SIGN" : message_hashed,
            "CB-ACCESS-TIMESTAMP" : time_current,
            "CB-ACCESS-PASSPHRASE" : api_struct.Password,
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "profile_id" : request_struct.Profile_id,
            "amount" : request_struct.Amount,
            "payment_method_id" : request_struct.Payment_method_id,
            "currency" : request_struct.Amount}).
        SetAuthToken(api_struct.Key).
        Post(api_struct.Host + request_struct.Request_path)

    // debug
    fmt.Println("Response Info:")
    fmt.Println("  Error      :", err)
    fmt.Println("  Status Code:", resp.StatusCode())
    fmt.Println("  Status     :", resp.Status())
    fmt.Println("  Proto      :", resp.Proto())
    fmt.Println("  Time       :", resp.Time())
    fmt.Println("  Received At:", resp.ReceivedAt())
    fmt.Println("  Body       :\n", resp)
    fmt.Println()

    return resp.StatusCode(), resp.Body()
}

func rest_post_transfer_coinbase(api_struct apiConfig, request_struct reqTransfer) (int, []byte) { //POST_REQUEST_WITHDRAW/DEPOSIT
    request_method := "POST"

    time_current := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message_hashed := gen_api_message(api_struct.Secret, time_current, request_method, request_struct.Request_path) //create hashed message to send

    client := resty.New() //create REST session

    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : api_struct.Key,
            "CB-ACCESS-SIGN" : message_hashed,
            "CB-ACCESS-TIMESTAMP" : time_current,
            "CB-ACCESS-PASSPHRASE" : api_struct.Password,
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "profile_id" : request_struct.Profile_id,
            "amount" : request_struct.Amount,
            "coinbase_account_id" : request_struct.Coinbase_account_id,
            "currency" : request_struct.Amount}).
        SetAuthToken(api_struct.Key).
        Post(api_struct.Host + request_struct.Request_path)

    // debug
    fmt.Println("Response Info:")
    fmt.Println("  Error      :", err)
    fmt.Println("  Status Code:", resp.StatusCode())
    fmt.Println("  Status     :", resp.Status())
    fmt.Println("  Proto      :", resp.Proto())
    fmt.Println("  Time       :", resp.Time())
    fmt.Println("  Received At:", resp.ReceivedAt())
    fmt.Println("  Body       :\n", resp)
    fmt.Println()

    return resp.StatusCode(), resp.Body()
}

func get_all_accounts(api_struct apiConfig) []apiAccount { //Get a list of trading accounts from the profile of the API key.
    request_path := "/accounts"

    var api_accounts []apiAccount

    response_status, response_body := rest_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_accounts); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("api_accounts:")
    fmt.Println()
    for account := range api_accounts {
        fmt.Println("api_accounts[", account, "]")
        fmt.Println(api_accounts[account].Id)
        fmt.Println(api_accounts[account].Currency)
        fmt.Println(api_accounts[account].Balance)
        fmt.Println(api_accounts[account].Hold)
        fmt.Println(api_accounts[account].Available)
        fmt.Println(api_accounts[account].Profile_id)
        fmt.Println(api_accounts[account].Trading_enabled)
        fmt.Println()
    }

    return api_accounts
}

func get_single_account(api_struct apiConfig, api_account_id string) apiAccount { //Information for a single account.
    request_path := "/accounts/" + api_account_id

    var api_account apiAccount //store single apiAccount

    response_status, response_body := rest_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_account); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("api_account:")
    fmt.Println(api_account.Id)
    fmt.Println(api_account.Currency)
    fmt.Println(api_account.Balance)
    fmt.Println(api_account.Hold)
    fmt.Println(api_account.Available)
    fmt.Println(api_account.Profile_id)
    fmt.Println(api_account.Trading_enabled)
    fmt.Println()

    return api_account
}

func get_single_account_holds(api_struct apiConfig, api_account_id string) []apiHold { //List the holds of an account that belong to the same profile as the API key.
    request_path := "/accounts/" + api_account_id + "/holds" //?limit=100" //implement limit logic later

    var api_account_holds []apiHold

    response_status, response_body := rest_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }
    if response_body == nil {
        return nil
    }

    if err := json.Unmarshal(response_body, &api_account_holds); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("api_account_holds:")
    for hold := range api_account_holds {
        fmt.Println("api_account_holds[", hold,"]")
        fmt.Println(api_account_holds[hold].Id)
        fmt.Println(api_account_holds[hold].Created_at)
        fmt.Println(api_account_holds[hold].Amount)
        fmt.Println(api_account_holds[hold].Ref)
        fmt.Println(api_account_holds[hold].Type)
        fmt.Println()
    }

    return api_account_holds
}

func get_single_account_ledgers(api_struct apiConfig, api_account_id string) []apiLedger { //List the holds of an account that belong to the same profile as the API key.
    request_path := "/accounts/" + api_account_id + "/ledger" //?limit=100" //implement limit logic later

    var api_account_ledgers []apiLedger

    response_status, response_body := rest_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_account_ledgers); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("account_id:", api_account_id)
    for ledger := range api_account_ledgers {
        fmt.Println(api_account_ledgers[ledger].Id)
        fmt.Println(api_account_ledgers[ledger].Amount)
        fmt.Println(api_account_ledgers[ledger].Balance)
        fmt.Println(api_account_ledgers[ledger].Created_at)
        fmt.Println(api_account_ledgers[ledger].Type)
        fmt.Println(api_account_ledgers[ledger].Details.Order_id)
        fmt.Println(api_account_ledgers[ledger].Details.Product_id)
        fmt.Println(api_account_ledgers[ledger].Details.Trade_id)
        fmt.Println()
    }

    return api_account_ledgers
}

func get_single_account_transfers(api_struct apiConfig, api_account_id string) []apiPastTransfer { //Lists past withdrawals and deposits for an account.
    request_path := "/accounts/" + api_account_id + "/transfers" //?limit=100" //implement limit logic later

    var api_account_transfers []apiPastTransfer

    response_status, response_body := rest_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_account_transfers); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("account_id:", api_account_id)
    for transfer := range api_account_transfers {
        fmt.Println(api_account_transfers[transfer].Id)
        fmt.Println(api_account_transfers[transfer].Type)
        fmt.Println(api_account_transfers[transfer].Created_at)
        fmt.Println(api_account_transfers[transfer].Completed_at)
        fmt.Println(api_account_transfers[transfer].Canceled_at)
        fmt.Println(api_account_transfers[transfer].Processed_at)
        fmt.Println(api_account_transfers[transfer].User_nonce)
        fmt.Println(api_account_transfers[transfer].Amount)
        fmt.Println(api_account_transfers[transfer].Details.Coinbase_payout_at)
        fmt.Println(api_account_transfers[transfer].Details.Coinbase_account_id)
        fmt.Println(api_account_transfers[transfer].Details.Coinbase_deposit_id)
        fmt.Println(api_account_transfers[transfer].Details.Coinbase_payment_method_id)
        fmt.Println(api_account_transfers[transfer].Details.Coinbase_payment_method_type)
        fmt.Println(api_account_transfers[transfer].Idem)
        fmt.Println()
    }

    return api_account_transfers 
}

func get_all_wallets(api_struct apiConfig) []apiWallet { //Gets all the user's available Coinbase wallets
    request_path := "/coinbase-accounts"

    var api_accounts_wallets []apiWallet

    response_status, response_body := rest_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_accounts_wallets); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("api_account_wallets: ", api_accounts_wallets)
    fmt.Println()
    for wallet := range api_accounts_wallets {
        fmt.Println()
        fmt.Println("api_accounts_wallets[", wallet, "]")
        fmt.Println(api_accounts_wallets[wallet].Id)
        fmt.Println(api_accounts_wallets[wallet].Name)
        fmt.Println(api_accounts_wallets[wallet].Balance)
        fmt.Println(api_accounts_wallets[wallet].Currency)
        fmt.Println(api_accounts_wallets[wallet].Type)
        fmt.Println(api_accounts_wallets[wallet].Primary)
        fmt.Println(api_accounts_wallets[wallet].Active)
        fmt.Println(api_accounts_wallets[wallet].Available_on_consumer)
        if api_accounts_wallets[wallet].Ready == (true || false) {
            fmt.Println(api_accounts_wallets[wallet].Ready)
        }
        if api_accounts_wallets[wallet].Wire_deposit_information.Account_name != "" {
            fmt.Println(api_accounts_wallets[wallet].Wire_deposit_information.Account_number)
            fmt.Println(api_accounts_wallets[wallet].Wire_deposit_information.Routing_number)
            fmt.Println(api_accounts_wallets[wallet].Wire_deposit_information.Bank_name)
            fmt.Println(api_accounts_wallets[wallet].Wire_deposit_information.Bank_address)
            fmt.Println(api_accounts_wallets[wallet].Wire_deposit_information.Bank_country.Code)
            fmt.Println(api_accounts_wallets[wallet].Wire_deposit_information.Bank_country.Name)
            fmt.Println(api_accounts_wallets[wallet].Wire_deposit_information.Account_name)
            fmt.Println(api_accounts_wallets[wallet].Wire_deposit_information.Account_address)
            fmt.Println(api_accounts_wallets[wallet].Wire_deposit_information.Reference)
        }
        if api_accounts_wallets[wallet].Swift_deposit_information.Account_name != "" {
            fmt.Println(api_accounts_wallets[wallet].Swift_deposit_information.Account_number)
            fmt.Println(api_accounts_wallets[wallet].Swift_deposit_information.Routing_number)
            fmt.Println(api_accounts_wallets[wallet].Swift_deposit_information.Bank_name)
            fmt.Println(api_accounts_wallets[wallet].Swift_deposit_information.Bank_address)
            fmt.Println(api_accounts_wallets[wallet].Swift_deposit_information.Bank_country.Code)
            fmt.Println(api_accounts_wallets[wallet].Swift_deposit_information.Bank_country.Name)
            fmt.Println(api_accounts_wallets[wallet].Swift_deposit_information.Account_name)
            fmt.Println(api_accounts_wallets[wallet].Swift_deposit_information.Account_address)
            fmt.Println(api_accounts_wallets[wallet].Swift_deposit_information.Reference)
        }
        if api_accounts_wallets[wallet].Sepa_deposit_information.Account_name != ""{
            fmt.Println(api_accounts_wallets[wallet].Sepa_deposit_information.Iban)
            fmt.Println(api_accounts_wallets[wallet].Sepa_deposit_information.Swift)
            fmt.Println(api_accounts_wallets[wallet].Sepa_deposit_information.Bank_name)
            fmt.Println(api_accounts_wallets[wallet].Sepa_deposit_information.Bank_address)
            fmt.Println(api_accounts_wallets[wallet].Sepa_deposit_information.Bank_country_name)
            fmt.Println(api_accounts_wallets[wallet].Sepa_deposit_information.Account_name)
            fmt.Println(api_accounts_wallets[wallet].Sepa_deposit_information.Account_address)
            fmt.Println(api_accounts_wallets[wallet].Sepa_deposit_information.Reference)
        }
        if api_accounts_wallets[wallet].Uk_deposit_information.Account_name != "" {
            fmt.Println(api_accounts_wallets[wallet].Uk_deposit_information.Sort_code)
            fmt.Println(api_accounts_wallets[wallet].Uk_deposit_information.Account_number)
            fmt.Println(api_accounts_wallets[wallet].Uk_deposit_information.Bank_name)
            fmt.Println(api_accounts_wallets[wallet].Uk_deposit_information.Account_name)
            fmt.Println(api_accounts_wallets[wallet].Uk_deposit_information.Reference)
        }
        if api_accounts_wallets[wallet].Destination_tag_name != "" {
            fmt.Println(api_accounts_wallets[wallet].Destination_tag_name)
        }
        if api_accounts_wallets[wallet].Destination_tag_regex != "" {
            fmt.Println(api_accounts_wallets[wallet].Destination_tag_regex)
        }
        fmt.Println(api_accounts_wallets[wallet].Hold_balance)
        fmt.Println(api_accounts_wallets[wallet].Hold_currency)
    }
    fmt.Println()
    
    return api_accounts_wallets
}

func generate_crypto_address(api_struct apiConfig, api_account_id string) apiCryptoAddress { //Generates a one-time crypto address for depositing crypto.
    request_path := "/coinbase-accounts/" + api_account_id + "/addresses"

    var api_account_address apiCryptoAddress

    response_status, response_body := rest_post_generate_address(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_account_address); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("api_account_address:", api_account_id)
    fmt.Println()
    fmt.Println(api_account_address.Id)
    fmt.Println(api_account_address.Address)
    if api_account_address.Address_info.Address != "" {
        fmt.Println(api_account_address.Address_info.Address)
        fmt.Println(api_account_address.Address_info.Destination_tag)
    }
    fmt.Println(api_account_address.Name)
    fmt.Println(api_account_address.Created_at)
    fmt.Println(api_account_address.Updated_at)
    if api_account_address.Network != "" {
        fmt.Println(api_account_address.Network)
    }
    if api_account_address.Uri_scheme != "" {
        fmt.Println(api_account_address.Uri_scheme)
    }
    fmt.Println(api_account_address.Resource)
    fmt.Println(api_account_address.Resource_path)
    if api_account_address.Warnings != nil {
        for warning := range api_account_address.Warnings {
            fmt.Println()
            fmt.Println("api_account_address.Warnings[", warning, "]")
            fmt.Println(api_account_address.Warnings[warning].Title)
            fmt.Println(api_account_address.Warnings[warning].Details)
            fmt.Println(api_account_address.Warnings[warning].Image_url)
        }
    }
    if api_account_address.Legacy_address != "" {
        fmt.Println(api_account_address.Legacy_address)
    }
    if api_account_address.Destination_tag != "" {
        fmt.Println(api_account_address.Destination_tag)
    }
    if api_account_address.Deposit_uri != "" {
        fmt.Println(api_account_address.Deposit_uri)
    }
    if api_account_address.Callback_url != "" {
        fmt.Println(api_account_address.Callback_url)
    }
    fmt.Println()

    return api_account_address
}

func convert_currency(api_struct apiConfig, profile_id string, from string, to string, amount string) apiConvert { //Converts funs from currency to currency
    request_path := "/conversions"

    var request_struct reqConvert
    var api_account_convert apiConvert

    request_struct.Request_path = request_path
    request_struct.Profile_id = profile_id
    request_struct.From = from
    request_struct.To = to
    request_struct.Amount = amount

    response_status, response_body := rest_post_convert_currency(api_struct, request_struct)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_account_convert); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    return api_account_convert
}

func deposit_coinbase_account(api_struct apiConfig, request_profile_id string, request_amount string, request_coinbase_account_id string, request_currency string) apiTransfer {
    request_path := "/deposits/coinbase-account"

    var request_struct reqTransfer
    var api_account_deposit apiTransfer

    request_struct.Request_path = request_path
    request_struct.Profile_id = request_profile_id
    request_struct.Amount = request_amount
    request_struct.Coinbase_account_id = request_coinbase_account_id
    request_struct.Currency = request_currency

    response_status, response_body := rest_post_transfer_coinbase(api_struct, request_struct)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_account_deposit); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    return api_account_deposit
}

func get_fees(api_struct apiConfig) apiFee {
    request_path := "/fees"

    var api_account_fees apiFee

    response_status, response_body := rest_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_account_fees); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("api_account_fees:")
    fmt.Println(api_account_fees.Taker_fee_rate)
    fmt.Println(api_account_fees.Maker_fee_rate)
    fmt.Println(api_account_fees.Usd_volume)
    fmt.Println()

    return api_account_fees
}

func get_all_fills(api_struct apiConfig, order_id string, product_id string, profile_id string, limit int64, before int64, after int64) []apiFill {
    request_path := "/fills"

    var api_account_fills []apiFill

    response_status, response_body := rest_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
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

func get_profiles(api_struct apiConfig, active bool) []apiProfile{
    request_path := "/profiles"

    var profiles []apiProfile

    response_status, response_body := rest_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
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

func get_signed_prices(api_struct apiConfig) apiPrice {
    request_path := "/oracle"

    var prices apiPrice

    response_status, response_body := rest_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
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

func rest_handler(api_struct apiConfig) {
    get_all_accounts(api_struct)
    get_signed_prices(api_struct)
}

func main() {
    f := "api.toml" //default configuration file
    if _, err := os.Stat(f); err != nil { //check configuration file exists
        fmt.Println("ERROR " + f + " does not exist")
        os.Exit(1)
    }

    var api_struct apiConfig //struct that stores API key information
    if _, err := toml.DecodeFile(f, &api_struct); err != nil { //decode TOML file
        fmt.Println("ERROR decoding toml configuration")
        os.Exit(1)
    }

    //add check for missing config elements
   
    rest_handler(api_struct) //rest handle function
}
