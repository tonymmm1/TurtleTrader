//Coinbase Pro API Websocket
package coinbase_pro

import (
    "encoding/json"
    "fmt"
    "os"
    "os/signal"

    "turtle/src/config"

    "github.com/gorilla/websocket"
)
const ( //Error messages (https://docs.cloud.coinbase.com/exchange/docs/websocket-overview)
    WEBSOCKET_ErrBufferFull int = 1
    WEBSOCKET_ErrSlowConsume = 2
    WEBSOCKET_ErrSlowRead = 3
)

//Message struct
type message struct { //Websocket request message
    Type string `json:"type"`
    Product_ids []string `json:"product_ids"`
    Channels []string `json:"channels"`
}

//Channels (https://docs.cloud.coinbase.com/exchange/docs/websocket-channels)
type messageHeartbeat struct { //The heartbeat channel
    Type string `json:"type"`
    Sequence string `json:"sequence"`
    Last_trade_id string `json:"last_trade_id"`
    Product_ids []string `json:"product_ids"`
    Time string `json:"time"`
}

type messageLevel2 struct { //The Level2 channel
    Type string `json:"type"`
    Product_id string `json:"product_id"`
    Changes_raw []interface{} `json:"changes"`
    Changes struct {
        Type string
        Price string
        Size string
    }
    Time string `json:"time"`
}

type messageTicker struct { //The Ticker channel
    Type string `json:"type"`
    Sequence string `json:"sequence"`
    Product_id string `json:"product_id"`
    Price string `json:"price"`
    Open_24h string `json:"open_24h"`
    Volume_24h string `json:"volume_24h"`
    Low_24h string `json:"low_24h"`
    High_24h string `json:"high_24h"`
    Volume_30d string `json:"volume_30d"`
    Best_bid string `json:"best_bid"`
    Best_ask string `json:"best_ask"`
    Side string `json:"side"`
    Time string `json:"time"`
    Trade_id string `json:"trade_id"`
    Last_size string `json:"last_size"`
}

//type messageUser struct {} //The user channel

type messageMatches struct { //The matches channel
    Type string `json:"type"`
    Trade_id string `json:"trade_id"`
    Maker_order_id string `json:"maker_order_id"`
    Taker_order_id string `json:"taker_order_id"`
    Side string `json:"side"`
    Size string `json:"size"`
    Price string `json:"price"`
    Product_id string `json:"product_id"`
    Sequence string `json:"sequence"`
    Time string `json:"time"`
}

type messageFull struct { //The full channel
    Order_id string `json:"order_id"`
    Reason string `json:"filled"`
    Price string `json:"price"`
    Remaining_size string `json:"remaining_size"`
    Type string `json:"type"`
    Side string `json:"side"`
    Product_id string `json:"product_id"`
    Time string `json:"time"`
    Sequence string `json:"sequence"`
}

func Websocket_run(product_ids []string, channels []string) { //Websocket handler
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)

    fmt.Println("connecting to", config.CBP.Websocket)

    c, _, err := websocket.DefaultDialer.Dial(config.CBP.Websocket, nil)
    if err != nil {
        fmt.Println("dial:", err)
    }
    defer c.Close()

    done := make(chan struct{})

    go func() {
        defer close(done)
        for {
            _, message, err := c.ReadMessage()
            if err != nil {
                fmt.Println("read:", err)
                return
            }

            fmt.Println(string(message))
        }
    }()

    message := message {
        Type: "subscribe",
        Product_ids: product_ids,
        Channels: channels,
    }

    message_json, err := json.Marshal(message)
    if err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    fmt.Println()
    fmt.Println("message:", string(message_json))
    fmt.Println()

    err = c.WriteMessage(websocket.TextMessage, message_json)
    if err != nil {
        fmt.Println("write:", err)
        return
    }

    for {
        select {
        case <-interrupt:
            fmt.Println("interrupt")

            // Cleanly close the connection by sending a close message and then
            // waiting (with timeout) for the server to close the connection.
            err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            if err != nil {
                fmt.Println("write close:", err)
                return
            }
            select {
            case <-done:
            }
            return
        }
    }
}
