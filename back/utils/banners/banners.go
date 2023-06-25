package banners

import (
	"log"
	"strings"
)

var boxArt = struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Right       string
	Left        string
	Top         string
	Bottom      string
}{
	TopLeft:     "┌",
	TopRight:    "┐",
	BottomLeft:  "└",
	BottomRight: "┘",
	Right:       "│",
	Left:        "│",
	Top:         "─",
	Bottom:      "─",
}

func padLine(line string) string {
	textLen := len(line)
	space := 80 - 4 - textLen
	leftSpace := space / 2
	rightSpace := space - leftSpace
	line = boxArt.Left + " " + strings.Repeat(" ", leftSpace) + line + strings.Repeat(" ", rightSpace) + " " + boxArt.Right
	return line
}

func BoxBanner(text string) string {

	maxLineLength := 80 - 4
	textLen := len(text)
	lines := []string{}

	for textLen > 0 {
		if textLen > maxLineLength {
			lines = append(lines, padLine(text[:maxLineLength]))
			text = text[maxLineLength:]
			textLen = len(text)
		} else {
			lines = append(lines, padLine(text))
			text = ""
		}
	}

	topBorder := boxArt.TopLeft + strings.Repeat(boxArt.Top, maxLineLength) + boxArt.TopRight
	bottomBorder := boxArt.BottomLeft + strings.Repeat(boxArt.Bottom, maxLineLength) + boxArt.BottomRight

	text = topBorder + "\n" + strings.Join(lines, "\n") + "\n" + bottomBorder

	return text
}

func Print(s string) {
	log.Print(BoxBanner(s))
}
