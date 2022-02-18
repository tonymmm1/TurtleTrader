//Coinbase Pro API : Orders
package coinbase_pro

import (
    "encoding/json"
    "fmt"
    "os"
)

type cbpFill struct { //Get all fills
    Trade_id int32 `json:"id"`
    Product_id string `json:"product_id"`
    Order_id string `json:"order_id"`
    User_id string `json:"user_id"`
    Profile_id string `json:"profile_id"`
    Liquidity string `json:"liquidity"`
    Price string `json:"price"`
    Size string `json:"size"`
    Fee string `json:"fee"`
    Created_at string `json:"created_at"`
    Side string `json:"side"`
    Settled bool `json:"settled"`
    Usd_volume string `json:"usd_volume"`
}

/*  Orders
*       Get all fills       (GET)
*       Get all orders      (GET)
*       Cancel all orders   (DELETE)
*       Create a new order  (POST)
*       Get single order    (GET)
*       Cancel an order     (DELETE)
*/

func cbp_get_all_fills(order_id string, product_id string, profile_id string, limit int64, before int64, after int64) []cbpFill {
    path := "/fills"

    var api_account_fills []cbpFill

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &api_account_fills); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get all fills")
    fmt.Println()
    for fill := range api_account_fills {
        fmt.Println("api_account_fills[", fill, "]")
        fmt.Println(api_account_fills[fill].Trade_id)
        fmt.Println(api_account_fills[fill].Product_id)
        fmt.Println(api_account_fills[fill].Order_id)
        fmt.Println(api_account_fills[fill].User_id)
        fmt.Println(api_account_fills[fill].Profile_id)
        fmt.Println(api_account_fills[fill].Liquidity)
        fmt.Println(api_account_fills[fill].Price)
        fmt.Println(api_account_fills[fill].Size)
        fmt.Println(api_account_fills[fill].Fee)
        fmt.Println()
    }

    return api_account_fills
}
