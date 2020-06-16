package termenv

import (
	"fmt"
	"strings"
)

const (
	AltScreenSeq             = "?1049h"
	ExitAltScreenSeq         = "?1049l"
	CursorUpSeq              = "%dA"
	CursorDownSeq            = "%dB"
	CursorForwardSeq         = "%dC"
	CursorBackSeq            = "%dD"
	CursorNextLineSeq        = "%dE"
	CursorPreviousLineSeq    = "%dF"
	CursorHorizontalSeq      = "%dG"
	CursorPositionSeq        = "%d;%dH"
	EraseDisplaySeq          = "%dJ"
	EraseLineSeq             = "%dK"
	ScrollUpSeq              = "%dS"
	ScrollDownSeq            = "%dT"
	SaveCursorPositionSeq    = "s"
	RestoreCursorPositionSeq = "u"
	ShowCursorSeq            = "?25h"
	HideCursorSeq            = "?25l"
	ChangeScrollingRegionSeq = "%d;%dr"
)

// Reset the terminal to its default style, removing any active styles.
func Reset() {
	fmt.Print(CSI + ResetSeq + "m")
}

// AltScreen switches to the altscreen. The former view can be restored with
// ExitAltScreen().
func AltScreen() {
	fmt.Print(CSI + AltScreenSeq)
}

// ExitAltScreen exits the altscreen and returns to the former terminal view.
func ExitAltScreen() {
	fmt.Print(CSI + ExitAltScreenSeq)
}

// ClearScreen clears the visible portion of the terminal.
func ClearScreen() {
	fmt.Printf(CSI+EraseDisplaySeq, 2)
	MoveCursor(1, 1)
}

// MoveCursor moves the cursor to a given position.
func MoveCursor(row int, column int) {
	fmt.Printf(CSI+CursorPositionSeq, row, column)
}

// HideCursor hides the cursor.
func HideCursor() {
	fmt.Printf(CSI + HideCursorSeq)
}

// ShowCursor shows the cursor.
func ShowCursor() {
	fmt.Printf(CSI + ShowCursorSeq)
}

// SaveCursorPosition saves the cursor position.
func SaveCursorPosition() {
	fmt.Print(CSI + SaveCursorPositionSeq)
}

// RestoreCursorPosition restores a saved cursor position.
func RestoreCursorPosition() {
	fmt.Print(CSI + RestoreCursorPositionSeq)
}

// CursorUp moves the cursor up a given number of lines.
func CursorUp(n int) {
	fmt.Printf(CSI+CursorUpSeq, n)
}

// CursorDown moves the cursor down a given number of lines.
func CursorDown(n int) {
	fmt.Printf(CSI+CursorDownSeq, n)
}

// CursorForward moves the cursor up a given number of lines.
func CursorForward(n int) {
	fmt.Printf(CSI+CursorForwardSeq, n)
}

// CursorBack moves the cursor backwards a given number of cells.
func CursorBack(n int) {
	fmt.Printf(CSI+CursorBackSeq, n)
}

// CursorNextLine moves the cursor down a given number of lines and places it at
// the beginning of the line.
func CursorNextLine(n int) {
	fmt.Printf(CSI+CursorNextLineSeq, n)
}

// CursorPrevLine moves the cursor up a given number of lines and places it at
// the beginning of the line.
func CursorPrevLine(n int) {
	fmt.Printf(CSI+CursorPreviousLineSeq, n)
}

// ClearLine clears the current line.
func ClearLine() {
	fmt.Printf(CSI+EraseLineSeq, 2)
}

// ClearLines clears a given number of lines.
func ClearLines(n int) {
	clearLine := fmt.Sprintf(CSI+EraseLineSeq, 2)
	cursorUp := fmt.Sprintf(CSI+CursorUpSeq, 1)
	fmt.Print(clearLine + strings.Repeat(cursorUp+clearLine, n))
}

// ChangeScrollingRegion sets the scrolling region of the terminal.
func ChangeScrollingRegion(top, bottom int) {
	fmt.Printf(CSI+ChangeScrollingRegionSeq, top, bottom)
}
