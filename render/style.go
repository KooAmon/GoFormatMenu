package render

import (
	"fmt"
	"strings"
)

type Style []Color

func New(colors ...Color) Style {
	return colors
}

func (s Style) Print(a ...interface{}) {
	printWithCode(s.String(), fmt.Sprint(a...))
}

func (s Style) String() string {
	return convertColorsToCode(s...)
}

func convertColorsToCode(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}

	var codes []string
	for _, color := range colors {
		codes = append(codes, color.String())
	}

	return strings.Join(codes, ";")
}

func printWithCode(code string, str string) {
	fmt.Print(RenderString(code, str))
}
