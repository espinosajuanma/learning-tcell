# Write on screen

## The stupid way

You can make a loop... But to understand:

- `row`
- `column`
. `Rune` to print
- `combining` - slice of runes
- `style`

```go
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
```

Another stupid way

```go
	s.SetContent(0, 0, 'H', []rune{'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'}, style)
```

## Why combining

I don't now yet. I'll try to understand what is the better way to print on a future. But this works.
