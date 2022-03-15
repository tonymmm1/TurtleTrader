package main

import (
    "fmt"
    "log"
    "time"

    //cbp "turtle/src/api/coinbase_pro"

    "github.com/jroimartin/gocui"
)

const (
    //headers
    VIEW_PRICES_HEADER = "prices_header"
    VIEW_TRADES_HEADER = "trades_header"
    VIEW_PROFITS_HEADER = "profits_header"
    VIEW_MANUAL_HEADER = "manual_header"
    VIEW_BALANCES_HEADER = "balances_header"

    //views
    VIEW_PRICES = "prices"
    VIEW_TRADES = "trades"
    VIEW_PROFITS = "profits"
    VIEW_MANUAL = "manual"
    VIEW_BALANCES = "balances"
)

func Ui_init() {
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()

    g.Highlight = true
    g.SelFgColor = gocui.ColorGreen
    g.BgColor = gocui.ColorBlack
    g.FgColor = gocui.ColorDefault

    g.SetManagerFunc(Ui_layout)

    if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Ui_quit); err != nil {
        log.Panicln(err)
    }

    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
}

func Ui_layout(g *gocui.Gui) error {
    maxX, maxY := g.Size()

    //VIEW_PRICES_HEADER
    if v, err := g.SetView(VIEW_PRICES_HEADER, 0, 0, maxX/6, 2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprintln(v, "Prices")
    }
    //VIEW_PRICES
    if v, err := g.SetView(VIEW_PRICES, 0, 2, maxX/6, maxY-2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprintln(v, "Prices")

        go func(g *gocui.Gui) {
            var count uint64 = 0
            for {
                //output := cbp.Get_all_currencies()
                g.Update(func(g *gocui.Gui) error {
                    v, err := g.View(VIEW_PRICES)
                        if err != nil {
                            return err
                        }

                        v.Clear()

                        //fmt.Fprint(v, output)
                        fmt.Fprintln(v, "refresh: 1 second")
                        fmt.Fprintln(v, "elapsed:", count, "seconds")
                        fmt.Fprintln(v, "time:", time.Now())
                        fmt.Fprintln(v, "unix:", time.Now().Unix())

                        count += 1
                        return nil 
                })
                time.Sleep(1 * time.Second)
            }

        }(g)
    }
    /*
    //VIEW_2
    if v, err := g.SetView(VIEW_2, maxX/2, 0, maxX/2+50, maxY-1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprintln(v, "Trades")
    }
    //VIEW_3
    if v, err := g.SetView(VIEW_3, maxX/2+50, 0, maxX/2+100, maxY-1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprintln(v, "Profit")
    }
    */
    return nil
    
}

func Ui_quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}
