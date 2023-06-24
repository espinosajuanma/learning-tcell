# First Lesson: Make it work

First lesson: `tcell` is more low level than I expected. But I don't mind. I kinda like it. It would be trickier to setup some stuff, but I can make my framework from scratch understanding how it works or just move to framework that use tcell under the hood later.

## Screen

I have to instantiate a screen and initialize it.

```go
func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
}
```

## Loop

Like any other "render" screen render we need to setup a loop. In the past I deal with this on some JavaScript frameworks with HTML 5 Canvas.

```go
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
```

## Events

In the previous loop there is a polling for events. I set it so `q` finish the execution of the program.

## Catch panics

This would need a better understanding. But they recommend to do it on the `tcell` tutorial.

```go
	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()
```
