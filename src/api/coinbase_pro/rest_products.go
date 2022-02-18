//Coinbase Pro API rest : Orders
package coinbase_pro

import (
        "fmt"
        "strconv"
        "time"

        "crypto-bot/src/config"

        "github.com/go-resty/resty/v2"
)

func rest_get_all_trading_pairs(path string, query_type string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    path2 := path + "?type=" + query_type

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
            "type" : query_type}).
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

func rest_get_product_book(path string, level int32) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    level2 := strconv.FormatInt(int64(level), 10)

    path2 := path + "?level=" + level2

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
            "level" : level2}).
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

func rest_get_product_candles(path string, granularity int32, start string, end string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    granularity2 := strconv.FormatInt(int64(granularity), 10)

    path2 := path + "?granularity=" + granularity2 + "&start=" + start + "&end=" + end

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
            "granularity" : granularity2,
            "start" : start,
            "end" : end}).
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

func rest_get_product_trades(path string, limit int32) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    limit2 := strconv.FormatInt(int64(limit), 10)

    path2 := path + "?limit=" + limit2

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
            "limit" : limit2}).
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
