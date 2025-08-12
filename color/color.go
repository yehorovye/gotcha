/*
simple package to print stuff with colors.
by yehorovye. published under CC0 1.0 (public domain)
*/

package color

import "os"

const (
	Reset = "\033[0m"

	// fg
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// bright fg
	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	// bg
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

// basic function to colorize & reset, isn't it great?
// note: this function will not colorize if "NO_COLOR"
// env is present. cheers bitch.
func Colorize(text, color string) string {
	if os.Getenv("NO_COLOR") == "1" {
		return text
	} else {
		return color + text + Reset
	}
}
