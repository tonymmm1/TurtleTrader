package coinbase_pro

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

func cbp_generate_message(time string, method string, path string, body string) string { //generate hashed message for REST requests
    message := time + method + path + body //construct prehase message

    decoded, err := base64.StdEncoding.DecodeString(CbpKey.Secret) //decode base64 encoded api secret
    if err != nil {
        fmt.Println("ERROR decoding api key secret")
        os.Exit(1)
    }

    hash := hmac.New(sha256.New, []byte(decoded)) //generate new SHA256 hmac based on decoded api secret
    hash.Write([]byte(message)) //hash message using hmac

    return base64.StdEncoding.EncodeToString(hash.Sum(nil)) //return hashed message
}

func cbp_rest_get(path string) (int, []byte) { //handles GET requests
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := cbp_generate_message(time, "GET", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetAuthToken(CbpKey.Key).
        Get(CbpKey.Host + path)

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

func cbp_rest_get_fee_estimate(path string, currency string, crypto_address string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    path2 := path + "?crypto_address=" + crypto_address + "&currency=" + currency

    message := cbp_generate_message(time, "GET", path2, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "currency" : currency,
            "crypto_address" : crypto_address}).
        SetAuthToken(CbpKey.Key).
        Get(CbpKey.Host + path)

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

func cbp_rest_get_fills(path string, order_id string, product_id string, profile_id string, limit int64, before int64, after int64) (int, []byte) { //handles GET requests
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := cbp_generate_message(time, "GET", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "order_id" : order_id,
            "product_id" : product_id,
            "profile_id" : profile_id,
            "limit" : strconv.FormatInt(limit, 10),
            "before" : strconv.FormatInt(before, 10),
            "after" : strconv.FormatInt(after, 10)}).
        SetAuthToken(CbpKey.Key).
        Get(CbpKey.Host + path)

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

func cbp_rest_get_all_trading_pairs(path string, query_type string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    path2 := path + "?type=" + query_type

    message := cbp_generate_message(time, "GET", path2, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "type" : query_type}).
        SetAuthToken(CbpKey.Key).
        Get(CbpKey.Host + path)

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

func cbp_rest_get_product_book(path string, level int32) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    level2 := strconv.FormatInt(int64(level), 10)

    path2 := path + "?level=" + level2

    message := cbp_generate_message(time, "GET", path2, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "level" : level2}).
        SetAuthToken(CbpKey.Key).
        Get(CbpKey.Host + path)

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

func cbp_rest_get_product_candles(path string, granularity int32, start string, end string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    granularity2 := strconv.FormatInt(int64(granularity), 10)

    path2 := path + "?granularity=" + granularity2 + "&start=" + start + "&end=" + end

    message := cbp_generate_message(time, "GET", path2, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "granularity" : granularity2,
            "start" : start,
            "end" : end}).
        SetAuthToken(CbpKey.Key).
        Get(CbpKey.Host + path)

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

func cbp_rest_get_product_trades(path string, limit int32) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    limit2 := strconv.FormatInt(int64(limit), 10)

    path2 := path + "?limit=" + limit2

    message := cbp_generate_message(time, "GET", path2, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "limit" : limit2}).
        SetAuthToken(CbpKey.Key).
        Get(CbpKey.Host + path)

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

func cbp_rest_get_profiles(path string, active bool) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := cbp_generate_message(time, "GET", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetQueryParams(map[string] string {
            "active" : strconv.FormatBool(active)}).
        SetAuthToken(CbpKey.Key).
        Get(CbpKey.Host + path)

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

func cbp_rest_post_create_profile(path string, name string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    name2 := "{\"name\":\"" + name + "\"}"
    
    message := cbp_generate_message(time, "POST", path, name2) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Accept" : "application/json",
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "name" : name}).
        SetAuthToken(CbpKey.Key).
        Post(CbpKey.Host + path)

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

func cbp_rest_post_transfer_funds_profiles(path string, from string, to string, currency string, amount string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    body := "{\"amount\":\"" + amount + "\","  +
            "\"currency\":\"" + currency + "\"," +
            "\"from\":\"" + from + "\"," +
            "\"to\":\"" + to + "\"}"

    message := cbp_generate_message(time, "POST", path, body) //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Accept" : "application/json",
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "from" : from,
            "to" : to,
            "currency" : currency,
            "amount" : amount}).
        SetAuthToken(CbpKey.Key).
        Post(CbpKey.Host + path)

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

func cbp_rest_post_address(path string) (int, []byte) { //POST_REQUEST_GENERATE_ADDRESS
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := cbp_generate_message(time, "POST", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetAuthToken(CbpKey.Key).
        Post(CbpKey.Host + path)

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

func cbp_rest_post_convert(path string, profile_id string, from string, to string, amount string, nonce string) (int, []byte) { //POST_REQUEST_CONVERT_CURRENCY
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := cbp_generate_message(time, "POST", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "profile_id" : profile_id,
            "from" : from,
            "to" : to,
            "amount" : amount,
            "nonce" : nonce}).
        SetAuthToken(CbpKey.Key).
        Post(CbpKey.Host + path)

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

func cbp_rest_get_convert(path string, profile_id string) (int, []byte) {
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := cbp_generate_message(time, "POST", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "profile_id" : profile_id}).
        SetAuthToken(CbpKey.Key).
        Post(CbpKey.Host + path)

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

func cbp_rest_post_coinbase(path string, profile_id string, amount string, coinbase_account_id string, currency string) (int, []byte) { //POST_REQUEST_WITHDRAW/DEPOSIT
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := cbp_generate_message(time, "POST", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "profile_id" : profile_id,
            "amount" : amount,
            "coinbase_account_id" : coinbase_account_id,
            "currency" : currency}).
        SetAuthToken(CbpKey.Key).
        Post(CbpKey.Host + path)

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

func cbp_rest_post_payment(path string, profile_id string, amount string, payment_method_id string, currency string) (int, []byte) { //POST_REQUEST_(WITHDRAW/DEPOSIT)_PAYMENT
    time := strconv.FormatInt(time.Now().Unix(), 10)    //store current Unix time as int

    message := cbp_generate_message(time, "POST", path, "") //create hashed message to send

    client := resty.New() //create REST session
    resp, err := client.R().
        SetHeader("Accept", "application/json").
        SetHeaders(map[string] string {
            "CB-ACCESS-KEY" : CbpKey.Key,
            "CB-ACCESS-SIGN" : message,
            "CB-ACCESS-TIMESTAMP" : time,
            "CB-ACCESS-PASSPHRASE" : CbpKey.Password,
            "Content-Type" : "application/json"}).
        SetBody(map[string] string {
            "profile_id" : profile_id,
            "amount" : amount,
            "payment_method_id" : payment_method_id,
            "currency" : currency}).
        SetAuthToken(CbpKey.Key).
        Post(CbpKey.Host + path)

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
