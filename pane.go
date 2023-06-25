package main

import "github.com/gdamore/tcell/v2"

type PaneMode string

const (
	PaneModeActive   = "active"
	PaneModeFocus    = "focus"
	PaneModeDisabled = "disabled"
)

var (
	DisabledPaneStyle tcell.Style = tcell.Style.Background(tcell.StyleDefault, tcell.ColorDefault).Foreground(tcell.ColorWhite)
	ActivePaneStyle   tcell.Style = tcell.Style.Background(tcell.StyleDefault, tcell.ColorDefault).Foreground(tcell.ColorGreen)
	FocusPaneStyle    tcell.Style = tcell.Style.Background(tcell.StyleDefault, tcell.ColorDefault).Foreground(tcell.ColorDarkMagenta)
)

type Pane struct {
	X      int
	Y      int
	Height int
	Width  int
	Screen *tcell.Screen
	Style  tcell.Style
	Mode   PaneMode
}

func (p *Program) NewPane(x, y, w, h int) *Pane {
	return &Pane{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
		Screen: &p.Screen,
		Style:  DisabledPaneStyle,
		Mode:   PaneModeDisabled,
	}
}

func (p *Pane) Focus() {
	p.Mode = PaneModeFocus
	p.Style = FocusPaneStyle
}

func (p *Pane) Active() {
	p.Mode = PaneModeFocus
	p.Style = ActivePaneStyle
}

func (p *Pane) Disable() {
	p.Mode = PaneModeActive
	p.Style = DisabledPaneStyle
}
