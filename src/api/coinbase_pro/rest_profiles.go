//Coinbase Pro API rest : Profiles
package coinbase_pro

import (
        "fmt"
        "strconv"
        "time"

        "crypto-bot/src/config"

        "github.com/go-resty/resty/v2"
)

func rest_get_profiles(path string, active bool) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := generate_message(time, "GET", path, "") //create hashed message to send

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
            "active" : strconv.FormatBool(active)}).
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

func rest_post_create_profile(path string, name string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    name2 := "{\"name\":\"" + name + "\"}"

    message := generate_message(time, "POST", path, name2) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Accept" : "application/json",
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "name" : name}).
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

func rest_post_transfer_funds_profiles(path string, from string, to string, currency string, amount string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    body := "{\"amount\":\"" + amount + "\","  +
            "\"currency\":\"" + currency + "\"," +
            "\"from\":\"" + from + "\"," +
            "\"to\":\"" + to + "\"}"

    message := generate_message(time, "POST", path, body) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Accept" : "application/json",
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "from" : from,
            "to" : to,
            "currency" : currency,
            "amount" : amount}).
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
