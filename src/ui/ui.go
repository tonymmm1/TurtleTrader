package ui

import (
    "fmt"
    "encoding/json"
    "log"
    "sort"
    "time"

    cbp "turtle/src/api/coinbase_pro"

    "github.com/jroimartin/gocui"
)

const (
    //views
    VIEW_PRICES = "prices"
    VIEW_TRADES = "trades"
    VIEW_PROFITS = "profits"
    VIEW_MANUAL = "manual"
    VIEW_BALANCES = "balances"
)

func Init() {
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()

    g.Highlight = true
    g.SelFgColor = gocui.ColorGreen
    g.BgColor = gocui.ColorBlack
    g.FgColor = gocui.ColorDefault

    g.SetManagerFunc(layout)

    if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
        log.Panicln(err)
    }

    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
}

func layout(g *gocui.Gui) error {
    maxX, maxY := g.Size()

    //VIEW_PRICES
    if _, err := g.SetView(VIEW_PRICES, 0, 0, maxX/6, maxY-1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }

        go func(g *gocui.Gui) { //redraw view
            for {
                g.Update(func(g *gocui.Gui) error {
                    v, err := g.View(VIEW_PRICES)
                        if err != nil {
                            return err
                        }

                        v.Clear()

                        v.Title = VIEW_PRICES

                        response := cbp.Get_signed_prices() //Get prices of cryptos

                        var prices cbp.Price

                        if err := json.Unmarshal(response, &prices); err != nil { //JSON unmarshal REST response body to store as struct
                            fmt.Println("ERROR decoding REST response")
                        }

                        //sort cryptos by name alphabetically
                        keys := make([]string, 0, len(prices.Prices))
                        for k := range prices.Prices {
                            keys = append(keys, k)
                        }
                        sort.Strings(keys)

                        for _, k := range keys {
                            fmt.Fprintln(v, k, prices.Prices[k])
                        }
                        
                        return nil 
                })
                time.Sleep(1 * time.Second) //sleep for 1 second
            }
        }(g)
    }
    //VIEW_TRADES
    if v, err := g.SetView(VIEW_TRADES, maxX/6 + 1, 0, maxX*3/4, 3 * maxY/5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Title = VIEW_TRADES
    }
    //VIEW_PROFITS
    if v, err := g.SetView(VIEW_PROFITS, maxX*3/4+1, 0, maxX - 1, maxY - 1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Title = VIEW_PROFITS
    }
    //VIEW_MANUAL
    if v, err := g.SetView(VIEW_MANUAL, maxX/6+1, maxY*3/5+1, maxX*3/4, maxY - 1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Title = VIEW_MANUAL
    }

    return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}
