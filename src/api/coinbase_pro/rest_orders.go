//Coinbase Pro API rest : Orders
package coinbase_pro

import (
        "fmt"
        "strconv"
        "strings"
        "time"

        "turtle/src/config"

        "github.com/go-resty/resty/v2"
)

func rest_get_fills(path string, order_id string, product_id string, profile_id string, limit int64, before int64, after int64) (int, []byte) { //handles GET requests
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := generate_message(time, "GET", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Accept" : "application/json",
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

func rest_get_all_orders(path string, limit int64, status []string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    limit2 := strconv.FormatInt(limit,10)
    status2 := strings.Join(status, "&status=") //build query param string to match resty generated string
    path2 := path + "?limit=" + limit2 + "&status=" + status2

    message := generate_message(time, "GET", path2, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Accept" : "application/json",
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "limit" : limit2}).
        SetQueryParamsFromValues(map[string] []string{
            "status" : status}).
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

func rest_delete_all_orders(path string, profile_id string, product_id string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    path2 := path + "?product_id=" + product_id + "&profile_id=" + profile_id

    message := generate_message(time, "DELETE", path2, "") //create hashed message to send 

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Accept" : "application/json",
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "profile_id" : profile_id,
            "product_id" : product_id}).
        SetAuthToken(config.CBP.Key).
        Delete(config.CBP.Host + path)

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
