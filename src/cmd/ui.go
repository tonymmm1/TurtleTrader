package main

import (
    "fmt"
    "log"
    "time"

    //cbp "turtle/src/api/coinbase_pro"

    "github.com/jroimartin/gocui"
)

const (
    VIEW_1 = "prices"
    VIEW_2 = "trades"
    VIEW_3 = "profit"
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

    //VIEW_1
    if v, err := g.SetView(VIEW_1, 0, 0, maxX/2, maxY-1); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }

        fmt.Fprintf(v, "Prices")

        go func(g *gocui.Gui) {
            var count uint64 = 0
            for {
                //output := cbp.Get_all_currencies()
                g.Update(func(g *gocui.Gui) error {
                    v, err := g.View(VIEW_1)
                        if err != nil {
                            return err
                        }

                        v.Clear()

                        fmt.Fprintln(v, "Prices")
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
    return nil
    
}

func Ui_quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}
