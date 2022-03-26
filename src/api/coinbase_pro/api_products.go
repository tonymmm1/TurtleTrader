//Coinbase Pro API : Products
package coinbase_pro

import (
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

func Get_all_trading_pairs(query_type string) []byte {
    path := "/products"

    status, response := rest_get_all_trading_pairs(path, query_type)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_product(product_id string) []byte {
    path := "/products/" + product_id

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_product_book(product_id string, level int32) []byte {
    path := "/products/" + product_id + "/book"

    status, response := rest_get_product_book(path, level)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_product_candles(product_id string, granularity int32, start string, end string) []byte {
    path := "/products/" + product_id + "/candles"

    status, response := rest_get_product_candles(path, granularity, start, end)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_product_stats(product_id string) []byte {
    path := "/products/" + product_id + "/stats"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return reponse
}

func Get_product_ticker(product_id string) []byte {
    path := "/products/" + product_id + "/ticker"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_product_trades(product_id string, limit int32) []byte {
    path := "/products/" + product_id + "/trades"

    status, response := rest_get_product_trades(path, limit)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
