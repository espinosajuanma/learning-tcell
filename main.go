package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func print(s tcell.Screen, text string, coord []int, style tcell.Style) {
	for i, c := range text {
		s.SetContent(coord[0]+i, coord[1], c, nil, style)
	}
}

func printMiddle(s tcell.Screen, text string) {
	style := tcell.Style.Background(tcell.StyleDefault, tcell.ColorGreen).Foreground(tcell.ColorBlack)
	w, h := s.Size()
	len := len(text)
	x := (w / 2) - (len / 2)
	y := h / 2
	for i, c := range text {
		s.SetContent(x+i, y, c, nil, style)
	}
}

func main() {
	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorGreen)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(style)
	s.Clear()
	printMiddle(s, "Hello World!")

	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Event loop
	for {
		s.Show() // Update screen

		ev := s.PollEvent()
		switch ev := ev.(type) { // Process event
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				return
			}
		}
	}
}
