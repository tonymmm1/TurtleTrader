package main

import (
        "crypto/hmac"
        "crypto/sha256"
        "encoding/base64"
        "fmt"
        "os"
        "strconv"
        "time"

        "github.com/go-resty/resty/v2"
)

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
