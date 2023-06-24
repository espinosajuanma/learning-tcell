package main

import "github.com/gdamore/tcell/v2"

type Mode string

const (
	NormalMode  = "normal"
	CommandMode = "command"
)

type Program struct {
	Mode    Mode
	Screen  tcell.Screen
	Message string
}

func NewProgram() (*Program, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return &Program{}, err
	}
	if err := s.Init(); err != nil {
		return &Program{}, err
	}
	p := Program{
		Mode:    NormalMode,
		Message: "Normal Mode",
		Screen:  s,
	}
	p.SetMode(NormalMode)

	return &p, nil
}

func (p *Program) Init() {
	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorGreen)

	p.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	p.Screen.SetStyle(style)
	p.Screen.Clear()
}

func (p *Program) SetMode(mode Mode) {
	p.Mode = mode
	switch mode {
	case NormalMode:
		p.Screen.HideCursor()
		p.Message = "Normal mode"
	case CommandMode:
		_, h := p.Screen.Size()
		p.Screen.ShowCursor(0, h-1)
		p.Message = "Command Mode"
	}
}

func (p *Program) Update() {
	for {
		p.PrintMode()
		p.Screen.Show()
		ev := p.Screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'q' && p.Mode == NormalMode {
				return
			} else if ev.Rune() == ':' && p.Mode == NormalMode {
				p.SetMode(CommandMode)
			} else if ev.Key() == tcell.KeyEsc && p.Mode == CommandMode {
				p.SetMode(NormalMode)
			}
		}
	}
}

func (p *Program) Quit() {
	maybePanic := recover()
	p.Screen.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}

func (p *Program) PrintMode() {
	style := tcell.StyleDefault
	w, h := p.Screen.Size()
	len := len(p.Mode)
	x := (w / 2) - (len / 2)
	y := h / 2
	p.Screen.Clear()
	for i, c := range p.Message {
		p.Screen.SetContent(x+i, y, c, nil, style)
	}
}
