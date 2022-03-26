//Coinbase Pro API : Orders
package coinbase_pro

import (
    "fmt"
    "os"
)

type Fill struct { //Get all fills
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

type Order struct { //Get single/all orders
    Id string `json:"id"`
    Price string `json:"price"`
    Size string `json:"size"`
    Product_id string `json:"product_id"`
    Profile_id string `json:"profile_id"`
    Side string `json:"side"`
    Funds string `json:"funds"`
    Specified_funds string `json:"specified_funds"`
    Type string `json:"type"`
    Time_in_force string `json:"time_in_force"`
    Expire_time string `json:"expire_time"`
    Post_only bool `json:"post_only"`
    Created_at string `json:"created_at"`
    Done_at string `json:"done_at"`
    Done_reason string `json:"done_reason"`
    Reject_reason string `json:"reject_reason"`
    Fill_fees string `json:"fill_fees"`
    Filled_size string `json:"filled_size"`
    Executed_value string `json:"executed_value"`
    Status string `json:"status"`
    Settled bool `json:"settled"`
    Stop string `json:"stop"`
    Stop_price string `json:"stop_price"`
    Funding_amount string `json:"funding_amount"`
    Client_oid string `json:"client_oid"`
}

const ( //Time in force https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postorders#time-in-force
    TIME_IN_FORCE_GTC string = "GTC"
    TIME_IN_FORCE_GTT = "GTT"
    TIME_IN_FORCE_IOC = "IOC"
    TIME_IN_FORCE_FOK = "FOK"
)

const ( //Self-trade prevention https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postorders#self-trade-prevention
    SELF_TRADE_PREV_DC string = "dc"
    SELF_TRADE_PREV_CO = "co"
    SELF_TRADE_PREV_CN = "cn"
    SELF_TRADE_PREV_CB = "cb"
)

/*  Orders
*       Get all fills       (GET)
*       Get all orders      (GET)
*       Cancel all orders   (DELETE)
*       Create a new order  (POST)
*       Get single order    (GET)
*       Cancel an order     (DELETE)
*/

func Get_all_fills(order_id string, product_id string, profile_id string, limit int64, before int64, after int64) []byte {
    path := "/fills"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_all_orders(limit int64, status []string) []byte {
    path := "/orders"

    status, response := rest_get_all_orders(path, limit, status)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Create_order_market_size(profile_id string, side string, product_id string, size float64) []byte {
    path := "/orders"

    status, response := rest_post_create_order_market_size(path, profile_id, side, product_id, size)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Create_order_market_fund(profile_id string, side string, product_id string, fund float64) []byte {
    path := "/orders"

    status, response := rest_post_create_order_market_fund(path, profile_id, side, product_id, fund)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Create_order_limit(profile_id string, side string, stp string, time_in_force string, cancel_after string, post_only bool, product_id string, price float64, size float64) []byte {
    path := "/orders"

    status, response := rest_post_create_order_limit(path, profile_id, side, stp, time_in_force, cancel_after, post_only, product_id, price, size)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Create_order_stop(profile_id string, side string, stp string, stop string, stop_price float64, time_in_force string, cancel_after string, product_id string, price float64, size float64) []byte {
    path := "/orders"

    status, response := rest_post_create_order_stop(path, profile_id, side, stp, stop, stop_price, time_in_force, cancel_after, product_id, price, size)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Cancel_all_orders(profile_id string, product_id string) []byte {
    path := "/orders"

    status, response := rest_delete_all_orders(path, profile_id, product_id)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_single_order(order_id string) []byte {
    path := "/orders/" + order_id

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Cancel_order(order_id string, profile_id string) []byte {
    path := "/orders/" + order_id

    status, response := rest_delete_order(path, order_id, profile_id)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
