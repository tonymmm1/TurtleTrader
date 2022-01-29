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

//Common return codes (https://docs.cloud.coinbase.com/exchange/docs/requests)
const (
    STATUS_CODE_SUCCESS int = 200           //Success
    STATUS_CODE_BAD_REQUEST = 400           //Bad Request -- Invalid request format
    STATUS_CODE_UNAUTHORIZED = 401          //Unauthorized -- Invalid API Key
    STATUS_CODE_FORBIDDEN = 403             //Forbidden -- You do not have access to the requested resource
    STATUS_CODE_NOT_FOUND = 404             //Not Found
    STATUS_CODE_INTERNAL_SERVER_ERROR = 500 //Internal Server Error -- We had a problem with our server
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

type apiTransfer struct { //struct to store API transfer
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

func rest_client_get(api_struct apiConfig, request_path string) (int, []byte) { //handles GET requests
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

    if resp == nil {
        return resp.StatusCode(), nil
    }

    return resp.StatusCode(), resp.Body()
}

/*
func rest_client_put() { //handles PUT requests
}

func rest_client_post() { //handles POST requests
}

func rest_client_delete() { //handles DELETE requests
}
*/

func get_all_accounts(api_struct apiConfig) []apiAccount { //Get a list of trading accounts from the profile of the API key.
    request_path := "/accounts"

    var api_accounts []apiAccount

    response_status, response_body := rest_client_get(api_struct, request_path)
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

    response_status, response_body := rest_client_get(api_struct, request_path)
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

    response_status, response_body := rest_client_get(api_struct, request_path)
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

    response_status, response_body := rest_client_get(api_struct, request_path)
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

func get_single_account_transfers(api_struct apiConfig, api_account_id string) []apiTransfer { //Lists past withdrawals and deposits for an account.
    request_path := "/accounts/" + api_account_id + "/transfers" //?limit=100" //implement limit logic later

    var api_account_transfers []apiTransfer

    response_status, response_body := rest_client_get(api_struct, request_path)
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

    response_status, response_body := rest_client_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_accounts_wallets); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("account_id:", api_accounts_wallets)
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
    
    return api_accounts_wallets
}

func rest_handler(api_struct apiConfig) {
    get_all_wallets(api_struct)
}

func main() {
    f := "api.toml" //default configuration file
    if _, err := os.Stat(f); err != nil { //check configuration file exists
        fmt.Println("ERROR " + f + " does not exist")
        os.Exit(1)
    }

    var api_struct apiConfig
    if _, err := toml.DecodeFile(f, &api_struct); err != nil { //decode TOML file
        fmt.Println("ERROR decoding toml configuration")
        os.Exit(1)
    }

    //add check for missing config elements
   
    rest_handler(api_struct)
}
