package main

import (
	"log"
)

func main() {
	p, err := NewProgram()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	p.Init()
	defer p.Quit()
	p.Update()
}
