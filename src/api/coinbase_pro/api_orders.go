//Coinbase Pro API : Orders
package coinbase_pro

import (
    "encoding/json"
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

/*  Orders
*       Get all fills       (GET)
*       Get all orders      (GET)
*       Cancel all orders   (DELETE)
*       Create a new order  (POST)
*       Get single order    (GET)
*       Cancel an order     (DELETE)
*/

func Get_all_fills(order_id string, product_id string, profile_id string, limit int64, before int64, after int64) []Fill {
    path := "/fills"

    var api_account_fills []Fill

    response_status, response_body := rest_get(path)
    if response_status != STATUS_CODE_SUCCESS {
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

func Get_all_orders(limit int64, status []string) []Order {
    path := "/orders"

    var orders []Order

    response_status, response_body := rest_get_all_orders(path, limit, status)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &orders); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get all orders")
    fmt.Println()
    for order := range orders {
        fmt.Println("orders[",order,"]")
        fmt.Println(orders[order].Id)
        fmt.Println(orders[order].Price)
        fmt.Println(orders[order].Size)
        fmt.Println(orders[order].Product_id)
        fmt.Println(orders[order].Profile_id)
        fmt.Println(orders[order].Side)
        fmt.Println(orders[order].Funds)
        fmt.Println(orders[order].Specified_funds)
        fmt.Println(orders[order].Type)
        fmt.Println(orders[order].Time_in_force)
        fmt.Println(orders[order].Expire_time)
        fmt.Println(orders[order].Post_only)
        fmt.Println(orders[order].Created_at)
        fmt.Println(orders[order].Done_at)
        fmt.Println(orders[order].Done_reason)
        fmt.Println(orders[order].Reject_reason)
        fmt.Println(orders[order].Fill_fees)
        fmt.Println(orders[order].Filled_size)
        fmt.Println(orders[order].Executed_value)
        fmt.Println(orders[order].Status)
        fmt.Println(orders[order].Settled)
        fmt.Println(orders[order].Stop)
        fmt.Println(orders[order].Stop_price)
        fmt.Println(orders[order].Funding_amount)
        fmt.Println(orders[order].Client_oid)
        fmt.Println()
    }

    return orders
}

//func Create_new_order()

func Cancel_all_orders(profile_id string, product_id string) []string {
    path := "/orders"

    var orders []string

    response_status, response_body := rest_delete_all_orders(path, profile_id, product_id)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &orders); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Cancel all orders")
    fmt.Println()
    for order := range orders {
        fmt.Println("order[",order,"]")
        fmt.Println(order)
        fmt.Println()
    }

    return orders
}

func Get_single_order(order_id string) Order {
    path := "/orders/" + order_id

    var order Order

    response_status, response_body := rest_get(path)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &order); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get single order")
    fmt.Println()
    fmt.Println(order.Id)
    fmt.Println(order.Price)
    fmt.Println(order.Size)
    fmt.Println(order.Product_id)
    fmt.Println(order.Profile_id)
    fmt.Println(order.Side)
    fmt.Println(order.Funds)
    fmt.Println(order.Specified_funds)
    fmt.Println(order.Type)
    fmt.Println(order.Time_in_force)
    fmt.Println(order.Expire_time)
    fmt.Println(order.Post_only)
    fmt.Println(order.Created_at)
    fmt.Println(order.Done_at)
    fmt.Println(order.Done_reason)
    fmt.Println(order.Reject_reason)
    fmt.Println(order.Fill_fees)
    fmt.Println(order.Filled_size)
    fmt.Println(order.Executed_value)
    fmt.Println(order.Status)
    fmt.Println(order.Settled)
    fmt.Println(order.Stop)
    fmt.Println(order.Stop_price)
    fmt.Println(order.Funding_amount)
    fmt.Println(order.Client_oid)
    fmt.Println()

    return order
}

func Cancel_order(order_id string, profile_id string) string {
    path := "/orders/" + order_id

    var order string

    response_status, response_body := rest_delete_order(path, order_id, profile_id)
    if response_status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &order); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Cancel an order")
    fmt.Println()
    fmt.Println(order)

    return order
}
