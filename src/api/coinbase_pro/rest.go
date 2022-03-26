//Coinbase Pro API rest
package coinbase_pro

import (
        "crypto/hmac"
        "crypto/sha256"
        "encoding/base64"
        "fmt"
        "os"
        "strconv"
        "time"

        "turtle/src/config"

        "github.com/go-resty/resty/v2"
)

func generate_message(time string, method string, path string, body string) string { //generate hashed message for REST requests
    message := time + method + path + body //construct prehase message

    decoded, err := base64.StdEncoding.DecodeString(config.CBP.Secret) //decode base64 encoded api secret
    if err != nil {
        fmt.Println("ERROR decoding api key secret")
        os.Exit(1)
    }

    hash := hmac.New(sha256.New, []byte(decoded)) //generate new SHA256 hmac based on decoded api secret
    hash.Write([]byte(message)) //hash message using hmac

    return base64.StdEncoding.EncodeToString(hash.Sum(nil)) //return hashed message
}

func rest_get(path string) (int, []byte) { //handles GET requests
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
