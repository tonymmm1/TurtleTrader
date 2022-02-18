//Coinbase Pro API : Products
package coinbase_pro

import (
    "encoding/json"
    "fmt"
    "os"
)

type TradingPair struct {
    Id string `json:"id"`
    Base_currency string `json:"base_currency"`
    Quote_currency string `json:"quote_currency"`
    Base_min_size string `json:"base_min_size"`
    Base_max_size string `json:"base_max_size"`
    Quote_increment string `json:"quote_increment"`
    Base_increment string `json:"base_increment"`
    Display_name string `json:"display_name"`
    Min_market_funds string `json:"min_market_funds"`
    Max_market_funds string `json:"max_market_funds"`
    Margin_enabled bool `json:"margin_enabled"`
    Post_only bool `json:"post_only"`
    Limit_only bool `json:limit_only"`
    Cancel_only bool `json:"cancel_only"`
    Status string `json:"status"`
    Status_message string `json:"status_message"`
    Trading_disabled bool `json:"trading_disabled"`
    Fx_stablecoin bool `json:"fx_stablecoin"`
    Max_slippage_percentage string `json:"max_slippage_percentage"`
    Auction_mode bool `json:"auction_mode"`
}

type ProductBook struct {
    Bids []interface{} `json:"bids"`
    Asks []interface{} `json:"asks"`
    Sequence float64 `json:"sequence"`
    Auction_mode bool `json:"auction_mode"`
    Auction struct {
        Open_price string `json:"open_price"`
        Open_size string `json:"open_size"`
        Best_bid_price string `json:"best_bid_price"`
        Best_bid_size string `json:"best_bid_size"`
        Best_ask_price string `json:"best_ask_price"`
        Best_ask_size string `json:"best_ask_size"`
        Auction_state string `json:"auction_state"`
        Can_open string `json:"can_open"`
        Time string `json:"time"`
    } `json:"auction"`
}

type ProductStats struct {
    Open string `json:"open"`
    High string `json:"high"`
    Low string `json:"low"`
    Last string `json:"last"`
    Volume string `json:"volume"`
    Volume_30day string `json:"volume_30day"`
}

type ProductTicker struct {
    Ask string `json:"ask"`
    Bid string `json:"bid"`
    Volume string `json:"volume"`
    Trade_id int32 `json:"trade_id"`
    Price string `json:"price"`
    Size string `json:"size"`
    Time string `json:"time"`
}

type ProductTrade struct {
    Trade_id int32 `json:"trade_id"`
    Side string `json:"side"`
    Size string `json:"size"`
    Price string `json:"price"`
    Time string `json:"time"`
}

/*  Products
*       Get all known trading pairs (GET)
*       Get single product          (GET)
*       Get product book            (GET)
*       Get product candles         (GET)
*       Get product stats           (GET)
*       Get product ticker          (GET)
*       Get product trades          (GET)
*/

func Get_all_trading_pairs(query_type string) []TradingPair {
    path := "/products"

    var trades []TradingPair

    response_status, response_body := rest_get_all_trading_pairs(path, query_type)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &trades); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get all known trading pairs")
    fmt.Println()
    for trade := range trades {
        fmt.Println("trades[", trade, "]")
        fmt.Println(trades[trade].Id)
        fmt.Println(trades[trade].Base_currency)
        fmt.Println(trades[trade].Quote_currency)
        fmt.Println(trades[trade].Base_min_size)
        fmt.Println(trades[trade].Base_max_size)
        fmt.Println(trades[trade].Quote_increment)
        fmt.Println(trades[trade].Base_increment)
        fmt.Println(trades[trade].Display_name)
        fmt.Println(trades[trade].Min_market_funds)
        fmt.Println(trades[trade].Max_market_funds)
        fmt.Println(trades[trade].Margin_enabled)
        fmt.Println(trades[trade].Post_only)
        fmt.Println(trades[trade].Limit_only)
        fmt.Println(trades[trade].Cancel_only)
        fmt.Println(trades[trade].Status)
        fmt.Println(trades[trade].Status_message)
        fmt.Println(trades[trade].Trading_disabled)
        fmt.Println(trades[trade].Fx_stablecoin)
        fmt.Println(trades[trade].Max_slippage_percentage)
        fmt.Println(trades[trade].Auction_mode)
        fmt.Println()
    }

    return trades
}

func Get_product(product_id string) TradingPair {
    path := "/products/" + product_id

    var product TradingPair

    response_status, response_body := rest_get(path)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &product); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get single product")
    fmt.Println()
    fmt.Println(product.Id)
    fmt.Println(product.Base_currency)
    fmt.Println(product.Quote_currency)
    fmt.Println(product.Base_min_size)
    fmt.Println(product.Base_max_size)
    fmt.Println(product.Quote_increment)
    fmt.Println(product.Base_increment)
    fmt.Println(product.Display_name)
    fmt.Println(product.Min_market_funds)
    fmt.Println(product.Max_market_funds)
    fmt.Println(product.Margin_enabled)
    fmt.Println(product.Post_only)
    fmt.Println(product.Limit_only)
    fmt.Println(product.Cancel_only)
    fmt.Println(product.Status)
    fmt.Println(product.Status_message)
    fmt.Println(product.Trading_disabled)
    fmt.Println(product.Fx_stablecoin)
    fmt.Println(product.Max_slippage_percentage)
    fmt.Println(product.Auction_mode)
    fmt.Println()

    return product
}

func Get_product_book(product_id string, level int32) ProductBook {
    path := "/products/" + product_id + "/book"

    var book ProductBook

    response_status, response_body := rest_get_product_book(path, level)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &book); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get product book")
    fmt.Println()

    for k, v := range book.Bids {
        fmt.Println("book.Bids[", k, "]")
        fmt.Println(k, ":", v)
        fmt.Println()
    }
    for k, v := range book.Asks {
        fmt.Println("book.Asks[", k, "]")
        fmt.Println(k, ":", v)
        fmt.Println()
    }
    fmt.Println()
    fmt.Println(book.Sequence)
    fmt.Println(book.Auction_mode)
    fmt.Println(book.Auction.Open_price)
    fmt.Println(book.Auction.Best_bid_price)
    fmt.Println(book.Auction.Best_bid_size)
    fmt.Println(book.Auction.Best_ask_price)
    fmt.Println(book.Auction.Best_ask_size)
    fmt.Println(book.Auction.Auction_state)
    fmt.Println(book.Auction.Can_open)
    fmt.Println(book.Auction.Time)
    fmt.Println()

    return book
}

func Get_product_candles(product_id string, granularity int32, start string, end string) []interface{} { //[]ProductCandle {
    path := "/products/" + product_id + "/candles"

    var candles []interface{}

    response_status, response_body := rest_get_product_candles(path, granularity, start, end)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &candles); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get product candles")
    fmt.Println()
    for k, v := range candles {
        fmt.Println(k, ":", v)
    }

    return candles
}

func Get_product_stats(product_id string) ProductStats {
    path := "/products/" + product_id + "/stats"

    var stats ProductStats

    response_status, response_body := rest_get(path)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &stats); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get product stats")
    fmt.Println()
    fmt.Println(stats.Open)
    fmt.Println(stats.High)
    fmt.Println(stats.Low)
    fmt.Println(stats.Last)
    fmt.Println(stats.Volume)
    fmt.Println(stats.Volume_30day)
    fmt.Println()

    return stats
}

func Get_product_ticker(product_id string) ProductTicker {
    path := "/products/" + product_id + "/ticker"

    var stats ProductTicker

    response_status, response_body := rest_get(path)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &stats); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get product ticker")
    fmt.Println()
    fmt.Println(stats.Ask)
    fmt.Println(stats.Bid)
    fmt.Println(stats.Volume)
    fmt.Println(stats.Trade_id)
    fmt.Println(stats.Price)
    fmt.Println(stats.Size)
    fmt.Println(stats.Time)
    fmt.Println()

    return stats
}

func Get_product_trades(product_id string, limit int32) []ProductTrade {
    path := "/products/" + product_id + "/trades"

    var trades []ProductTrade

    response_status, response_body := rest_get_product_trades(path, limit)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &trades); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get product trades")
    fmt.Println()
    for trade := range trades {
        fmt.Println("trades[", trade, "]")
        fmt.Println(trades[trade].Trade_id)
        fmt.Println(trades[trade].Side)
        fmt.Println(trades[trade].Size)
        fmt.Println(trades[trade].Price)
        fmt.Println(trades[trade].Time)
        fmt.Println()
    }

    return trades
}
