//Coinbase Pro API rest : Orders
package coinbase_pro

import (
        "fmt"
        "strconv"
        "time"

        "crypto-bot/src/config"

        "github.com/go-resty/resty/v2"
)

func rest_get_fills(path string, order_id string, product_id string, profile_id string, limit int64, before int64, after int64) (int, []byte) { //handles GET requests
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
            "order_id" : order_id,
            "product_id" : product_id,
            "profile_id" : profile_id,
            "limit" : strconv.FormatInt(limit, 10),
            "before" : strconv.FormatInt(before, 10),
            "after" : strconv.FormatInt(after, 10)}).
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
