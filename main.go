package main

import (
        "crypto/hmac"
        "crypto/sha256"
        "encoding/base64"
        //"encoding/json"
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

type tomlConfig struct { //configuration toml file struct
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

func gen_api_message(api_key_password string, api_key_secret string, time_current string, request_method string, request_path string) string {
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

func rest_client_get(api_host string, api_key string, api_key_password string, api_key_secret string, request_path string) (int, []byte) { //handles GET requests
    request_method := "GET"
    time_current := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message_hashed := gen_api_message(api_key_password, api_key_secret, time_current, request_method, request_path) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : api_key,
            "CB-ACCESS-SIGN" : message_hashed,
            "CB-ACCESS-TIMESTAMP" : time_current,
            "CB-ACCESS-PASSPHRASE" : api_key_password,
            "Content-Type" : "application/json"}).
        SetAuthToken(api_key).
        Get(api_host + request_path)
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

/*
func rest_client_put() { //handles PUT requests
}

func rest_client_post() { //handles POST requests
}

func rest_client_delete() { //handles DELETE requests
}
*/

func get_all_accounts(api_host string, api_key string, api_key_password string, api_key_secret string) int { //handles GET requests
    request_path := "/accounts"

    resp,_ := rest_client_get(api_host, api_key, api_key_password, api_key_secret, request_path)
    if resp != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", resp)
        os.Exit(1)
    }

    return resp
}

func rest_handler(api_host string, api_key string, api_key_password string, api_key_secret string) {
    get_all_accounts(api_host, api_key, api_key_password, api_key_secret)
}

func main() {
    f := "api.toml" //default configuration file
    if _, err := os.Stat(f); err != nil { //check configuration file exists
        fmt.Println("ERROR " + f + " does not exist")
        os.Exit(1)
    }

    var config tomlConfig
    if _, err := toml.DecodeFile(f, &config); err != nil { //decode TOML file
        fmt.Println("ERROR decoding toml configuration")
        os.Exit(1)
    }
   
    //store values from configuration file
    api_host := config.Host
    api_key := config.Key
    api_key_password := config.Password
    api_key_secret := config.Secret

    rest_handler(api_host, api_key, api_key_password, api_key_secret)
}
