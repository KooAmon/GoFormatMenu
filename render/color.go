package render

import "strconv"

type Color uint8

const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

const (
	FgGrey Color = iota + 90
	FgBrightRed
	FgBrightGreen
	FgBrightYellow
	FgBrightBlue
	FgBrightMagenta
	FgBrightCyan
	FgBrightWhite
)

const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

const (
	BgGrey Color = iota + 100
	BgBrightRed
	BgBrightGreen
	BgBrightYellow
	BgBrightBlue
	BgBrightMagenta
	BgBrightCyan
	BgBrightWhite
)

var FgColors = map[string]Color{
	"Black":         FgBlack,
	"Red":           FgRed,
	"Green":         FgGreen,
	"Yellow":        FgYellow,
	"Blue":          FgBlue,
	"Magenta":       FgMagenta,
	"Cyan":          FgCyan,
	"White":         FgWhite,
	"Grey":          FgGrey,
	"BrightRed":     FgBrightRed,
	"BrightGreen":   FgBrightGreen,
	"BrightYellow":  FgBrightYellow,
	"BrightBlue":    FgBrightBlue,
	"BrightMagenta": FgBrightMagenta,
	"BrightCyan":    FgBrightCyan,
	"BrightWhite":   FgBrightWhite,
}

var BgColors = map[string]Color{
	"Black":         BgBlack,
	"Red":           BgRed,
	"Green":         BgGreen,
	"Yellow":        BgYellow,
	"Blue":          BgBlue,
	"Magenta":       BgMagenta,
	"Cyan":          BgCyan,
	"White":         BgWhite,
	"Grey":          BgGrey,
	"BrightRed":     BgBrightRed,
	"BrightGreen":   BgBrightGreen,
	"BrightYellow":  BgBrightYellow,
	"BrightBlue":    BgBrightBlue,
	"BrightMagenta": BgBrightMagenta,
	"BrightCyan":    BgBrightCyan,
	"BrightWhite":   BgBrightWhite,
}

func (c Color) String() string {
	return strconv.FormatInt(int64(c), 10)
}
