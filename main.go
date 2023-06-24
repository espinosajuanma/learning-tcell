package main

import (
	"fmt"

	"github.com/gdamore/encoding"
	"github.com/gdamore/tcell/v2"
)

func main() {
	// Just throwing some random stuff...
	// If it compiles is fine, right? :3
	c := tcell.PaletteColor(0)
	fmt.Println(c)
	tcell.RegisterEncoding("utf-8", encoding.UTF8)
	tcell.NewConsoleScreen()
	tcell.NewEventKey(tcell.KeyCtrlC, 'q', tcell.ModCtrl)
	tcell.NewScreen()
	tcell.NewEventError(fmt.Errorf("Hasdsad"))
}
