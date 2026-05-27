// Package output
package output

const (
	reset = "\033[0m"

	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	cyan   = "\033[36m"
	white  = "\033[37m"

	bold = "\033[1m"

	boldRed    = "\033[1;31m"
	boldGreen  = "\033[1;32m"
	boldYellow = "\033[1;33m"
)

// Generic

func Colorize(color, text string) string { return color + text + reset }

// Helper
// Normal colors

func RedColor(text string) string    { return Colorize(red, text) }
func GreenColor(text string) string  { return Colorize(green, text) }
func YellowColor(text string) string { return Colorize(yellow, text) }
func BlueColor(text string) string   { return Colorize(blue, text) }
func CyanColor(text string) string   { return Colorize(cyan, text) }
func WhiteColor(text string) string  { return Colorize(white, text) }

// Bold colors

func BoldColor(text string) string       { return Colorize(bold, text) }
func BoldRedColor(text string) string    { return Colorize(boldRed, text) }
func BoldGreenColor(text string) string  { return Colorize(boldGreen, text) }
func BoldYellowColor(text string) string { return Colorize(boldYellow, text) }
