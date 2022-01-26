package main

import (
        "crypto/hmac"
        "crypto/sha256"
        "encoding/base64"
        "fmt"
        "strconv"
        "time"

        "github.com/go-resty/resty/v2"
)

func main(){

    api_host := "https://api-public.sandbox.exchange.coinbase.com"
    api_key := ""
    api_key_password := ""
    api_key_secret := ""
    request_method := "GET"
    request_path := "/accounts"

    time_current := time.Now().Unix() //1000 //time in ms

    message := strconv.FormatInt(time_current, 10) + request_method + request_path //construct prehase message

    fmt.Println("debug> raw message ", message)

    decoded_secret, err := base64.StdEncoding.DecodeString(api_key_secret) //decode base64 encoded api secret
    if err != nil {
        fmt.Println("ERROR decoding api key secret")
    }

    fmt.Println("debug> decoded api secret: ", decoded_secret) //debug

    hash := hmac.New(sha256.New, []byte(decoded_secret))
    hash.Write([]byte(message))

    fmt.Println("debug> hmac: ", hash) //debug

    message_hashed := base64.StdEncoding.EncodeToString(hash.Sum(nil)) 
    
    fmt.Println("debug> hashed message ", message_hashed) //debug

    client := resty.New()
    resp, err := client.R().
    SetHeader("Accept", "application/json"). 
    SetHeaders(map[string] string {"CB-ACCESS-KEY" : api_key, "CB-ACCESS-SIGN" : message_hashed, "CB-ACCESS-TIMESTAMP" : strconv.FormatInt(time_current, 10), 
    "CB-ACCESS-PASSPHRASE" : api_key_password, "Content-Type" : "application/json"}).
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
