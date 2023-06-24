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

func printHelloWorld(s tcell.Screen, style tcell.Style) {
	// This is stupid, but is clearer to understand than a loop

	s.SetContent(0, 0, 'H', nil, style)
	s.SetContent(1, 0, 'e', nil, style)
	s.SetContent(2, 0, 'l', nil, style)
	s.SetContent(3, 0, 'l', nil, style)
	s.SetContent(4, 0, 'o', nil, style)
	s.SetContent(5, 0, ' ', nil, style)
	s.SetContent(6, 0, 'W', nil, style)
	s.SetContent(7, 0, 'o', nil, style)
	s.SetContent(8, 0, 'r', nil, style)
	s.SetContent(9, 0, 'l', nil, style)
	s.SetContent(10, 0, 'd', nil, style)
	s.SetContent(11, 0, '!', nil, style)

	s.SetContent(0, 1, 'H', []rune{'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'}, style)
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
	printHelloWorld(s, style)
	print(s, "Hello World!", []int{0, 2}, style)

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
