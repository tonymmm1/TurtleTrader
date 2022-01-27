package main

import (
        "crypto/hmac"
        "crypto/sha256"
        "encoding/base64"
        "fmt"
        "os"
        "strconv"
        "time"

        "github.com/BurntSushi/toml"
        "github.com/go-resty/resty/v2"
)

type tomlConfig struct { //configuration toml file struct
    Host string 
    Key string
    Password string
    Secret string
}

func gen_api_message(api_key_password string, api_key_secret string, time_current string, request_method string, request_path string) string{

    message := time_current + request_method + request_path //construct prehase message

    decoded_secret, err := base64.StdEncoding.DecodeString(api_key_secret) //decode base64 encoded api secret
    if err != nil {
        fmt.Println("ERROR decoding api key secret")
        os.Exit(1)
    }

    hash := hmac.New(sha256.New, []byte(decoded_secret))
    hash.Write([]byte(message))

    return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func main() {
    f := "api.toml" //default configuration

    if _, err := os.Stat(f); err != nil { //check configuration file exists
        fmt.Println("ERROR " + f + " does not exist")
        os.Exit(1)
    }

    var config tomlConfig
    _, err := toml.DecodeFile(f, &config)
    if err != nil {
        fmt.Println("ERROR decoding toml configuration")
        os.Exit(1)
    }
   
    //store values from configuration file
    api_host := config.Host
    api_key := config.Key
    api_key_password := config.Password
    api_key_secret := config.Secret

    request_method := "GET"
    request_path := "/accounts"

    //store current Unix time as int
    time_current := strconv.FormatInt(time.Now().Unix(), 10)

    //create hashed message to send 
    message_hashed := gen_api_message(api_key_password, api_key_secret, time_current, request_method, request_path)

    //REST client
    client := resty.New()
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

    fmt.Println("Response Info:")
    fmt.Println("  Error      :", err)
    fmt.Println("  Status Code:", resp.StatusCode())
    fmt.Println("  Status     :", resp.Status())
    fmt.Println("  Proto      :", resp.Proto())
    fmt.Println("  Time       :", resp.Time())
    fmt.Println("  Received At:", resp.ReceivedAt())
    fmt.Println("  Body       :\n", resp)
    fmt.Println()
}
