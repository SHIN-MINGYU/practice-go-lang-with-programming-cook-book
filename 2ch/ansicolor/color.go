package ansicolor

import "fmt"

// Text Color
type Color int

const (
	// default color
	ColorNone = iota

	// red text
	Red 
	//green text
	Green 
    // yellow text
    Yellow 
    // blue text
    Blue 
    // magenta text
    Magenta 
    // cyan text
    Cyan 
    // white text
    White
	// black text
	Black Color = -1
)

// ColorText structure save color values and color 
type ColorText struct {
	TextColor Color
	Text      string
}

func (c *ColorText) String() string {
	if c.TextColor == ColorNone {
		return c.Text
	} 
	value := 30
	if c.TextColor != Black {
		value += int(c.TextColor)
	}

	return fmt.Sprintf("\033[0;%dm%s\003[0m",value, c.Text)
}