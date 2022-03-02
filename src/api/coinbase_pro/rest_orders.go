//Coinbase Pro API rest : Orders
package coinbase_pro

import (
        "encoding/json"
        "fmt"
        "strconv"
        "strings"
        "time"
        "os"

        "turtle/src/config"

        "github.com/go-resty/resty/v2"
)

type restOrdersMarketBodySize struct {
    Profile_id string `json:"profile_id"`
    Type string `json:"type"`
    Side string `json:"side"`
    Product_id string `json:"product_id"`
    Size string `json:"size"`
}

type restOrdersMarketBodyPrice struct {
    Profile_id string `json:"profile_id"`
    Type string `json:"type"`
    Side string `json:"side"`
    Product_id string `json:"product_id"`
    Funds string `json:"funds"`
}

type restOrdersStopBody struct {
    Profile_id string `json:"profile_id"`
    Type string `json:"type"`
    Side string `json:"side"`
    Stp string `json:"stp"`
    Stop string `json:"stop"`
    Stop_price string `json:"stop_price"`
    Time_in_force string `json:"time_in_force"`
    Cancel_after string `json:"cancel_after"`
    Product_id string `json:"product_id"`
    Price string `json:"price"`
    Size string `json:"size"`
}

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

func rest_post_create_order_market_size(path string, profile_id string, side string, product_id string, size float64) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    body := restOrdersMarketBodySize {
        Profile_id: profile_id,
        Type: "market",
        Side: side,
        Product_id: product_id,
        Size : fmt.Sprintf("%f", size),
    }

    body_json, err := json.Marshal(body)
    if err != nil {
        fmt.Println("ERROR encoding JSON string")
        os.Exit(1)
    }

    message := generate_message(time, "POST", path, string(body_json)) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Accept" : "application/json",
            "Content-Type" : "application/json"}).
        SetBody(string(body_json)).
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

func rest_post_create_order_market_fund(path string, profile_id string, side string, product_id string, fund float64) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    body := restOrdersMarketBodyPrice {
        Profile_id: profile_id,
        Type: "market",
        Side: side,
        Product_id: product_id,
        Funds: fmt.Sprintf("%f", fund),
    }

    body_json, err := json.Marshal(body)
    if err != nil {
        fmt.Println("ERROR encoding JSON string")
        os.Exit(1)
    }

    message := generate_message(time, "POST", path, string(body_json)) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Accept" : "application/json",
            "Content-Type" : "application/json"}).
        SetBody(string(body_json)).
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

func rest_post_create_order_stop(path string, profile_id string, side string, stp string, stop string, stop_price float64, time_in_force string, cancel_after string, product_id string, price float64, size float64) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    body := restOrdersStopBody {
        Profile_id: profile_id,
        Type: "limit",
        Side: side,
        Stp: stp,
        Stop: stop,
        Stop_price: strconv.FormatFloat(stop_price, 'f', -1, 64),
        Time_in_force: time_in_force,
        Cancel_after: cancel_after,
        Product_id: product_id,
        Price: strconv.FormatFloat(price, 'f', -1, 64),
        Size: strconv.FormatFloat(size, 'f', -1, 64),
    }

    body_json, err := json.Marshal(body)
    if err != nil {
        fmt.Println("ERROR encoding JSON string")
        os.Exit(1)
    }

    message := generate_message(time, "POST", path, string(body_json)) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : config.CBP.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : config.CBP.Password,
            "Accept" : "application/json",
            "Content-Type" : "application/json"}).
        SetBody(string(body_json)).
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

func rest_delete_order(path string, order_id string, profile_id string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    path2 := path + "?order_id=" + order_id + "&profile_id=" + profile_id

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
            "order_id" : order_id,
            "profile_id" : profile_id}).
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
