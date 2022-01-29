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

type apiAccount struct { //array to store API account struct
    Id string `json:"id"`
    Currency string `json:"currency"`
    Balance string `json:"balance"`
    Hold string `json:"hold"`
    Available string `json:"available"`
    Profile_id string `json:"profile_id"`
    Trading_enabled bool `json:"trading_enabled"`
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
    fmt.Println()
    for api_account := range api_accounts {
        fmt.Printf("api_accounts[%d]:\n", api_account)
        fmt.Println(api_accounts[api_account].Id)
        fmt.Println(api_accounts[api_account].Currency)
        fmt.Println(api_accounts[api_account].Balance)
        fmt.Println(api_accounts[api_account].Hold)
        fmt.Println(api_accounts[api_account].Available)
        fmt.Println(api_accounts[api_account].Profile_id)
        fmt.Println(api_accounts[api_account].Trading_enabled)
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

func get_single_account_holds(api_struct apiConfig, api_account_id string) []byte { //List the holds of an account that belong to the same profile as the API key.
    request_path := "/accounts/" + api_account_id + "/holds" //?limit=100" //implement limit logic later

    response_status, response_body := rest_client_get(api_struct, request_path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }
    if response_body != nil {
        return response_body
    }

    return nil
}

func rest_handler(api_struct apiConfig) {
    get_all_accounts(api_struct)
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
