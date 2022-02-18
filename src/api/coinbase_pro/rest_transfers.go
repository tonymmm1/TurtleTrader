//Coinbase Pro API rest : Transfers
package coinbase_pro

import (
        "fmt"
        "strconv"
        "time"

        "crypto-bot/src/config"

        "github.com/go-resty/resty/v2"
)

func rest_post_coinbase(path string, profile_id string, amount string, coinbase_account_id string, currency string) (int, []byte) { //POST_REQUEST_WITHDRAW/DEPOSIT
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := generate_message(time, "POST", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "profile_id" : profile_id,
            "amount" : amount,
            "coinbase_account_id" : coinbase_account_id,
            "currency" : currency}).
        SetAuthToken(config.CBP.Key).
        Post(config.CBP.Host + path)

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

func rest_post_payment(path string, profile_id string, amount string, payment_method_id string, currency string) (int, []byte) { //POST_REQUEST_(WITHDRAW/DEPOSIT)_PAYMENT
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := generate_message(time, "POST", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "profile_id" : profile_id,
            "amount" : amount,
            "payment_method_id" : payment_method_id,
            "currency" : currency}).
        SetAuthToken(config.CBP.Key).
        Post(config.CBP.Host + path)

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

func rest_get_fee_estimate(path string, currency string, crypto_address string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    path2 := path + "?crypto_address=" + crypto_address + "&currency=" + currency

    message := generate_message(time, "GET", path2, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "currency" : currency,
            "crypto_address" : crypto_address}).
        SetAuthToken(config.CBP.Key).
        Get(config.CBP.Host + path)

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
