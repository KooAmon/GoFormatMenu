package menu

import (
	"GoFormatMenu/models"
	"GoFormatMenu/render"
	"strconv"
	"strings"
)

var (
	CharsFront  string
	CharsBack   string
	MenuStyle   render.Style
	BorderStyle render.Style
	MenuFG      render.Color
	MenuBG      render.Color
	BorderFG    render.Color
	BorderBG    render.Color
)

func WriteMenu(menu *models.Menu) int {

	render.ClearScreen()
	assignColorValues(menu.Colors)
	assignStyle()
	TotalMenuItems := len(menu.MenuItems) - 1

	FrontBuffer := len(strconv.Itoa(TotalMenuItems)) + 2
	BackBuffer := menu.LineWidth - FrontBuffer - 3
	CharsFront = strings.Repeat(HorizontalBorder, FrontBuffer)
	CharsBack = strings.Repeat(HorizontalBorder, BackBuffer)

	writeTopBorder()
	writeHeader(menu.Header, &FrontBuffer, &BackBuffer)
	writeMiddleBorder()

	for index, menuItem := range menu.MenuItems {
		writeMenuItem(menuItem.Key, index, &FrontBuffer, &BackBuffer)
		index++
	}
	writeBottomBorder()

	return *ValidateUserInputIsInteger(len(menu.MenuItems))
}

func writeHeader(header string, frontBuffer *int, backBuffer *int) {
	BackCharBuffer := strings.Repeat(Space, (*backBuffer - len(header) - 1))
	FrontCharBuffer := strings.Repeat(Space, *frontBuffer)

	BorderStyle.Print(VerticalBorder)
	MenuStyle.Print(FrontCharBuffer)
	BorderStyle.Print(MiddleBorder)
	MenuStyle.Print(Space, header, BackCharBuffer)
	BorderStyle.Print(VerticalBorder, "\n")
}

func writeMenuItem(menuItem string, index int, frontBuffer *int, backBuffer *int) {

	FrontCharBuffer := strings.Repeat(Space, (*frontBuffer - len(strconv.Itoa(index)) - 1))
	BackCharBuffer := strings.Repeat(Space, (*backBuffer - len(menuItem) - 1))

	BorderStyle.Print(VerticalBorder)
	MenuStyle.Print(FrontCharBuffer, index, Space)
	BorderStyle.Print(MiddleBorder)
	MenuStyle.Print(Space, menuItem, BackCharBuffer)
	BorderStyle.Print(VerticalBorder, "\n")
}

func writeTopBorder() {
	BorderStyle.Print(TopLeftCorner, CharsFront, TopMiddleIntersection, CharsBack, TopRightCorner, "\n")
}

func writeMiddleBorder() {
	BorderStyle.Print(MiddleLeftIntersection, CharsFront, MiddleMiddleIntersection, CharsBack, MiddleRightIntersection, "\n")
}

func writeBottomBorder() {
	BorderStyle.Print(BottomLeftCorner, CharsFront, BottomMiddleIntersection, CharsBack, BottomRightCorner, "\n")
}

func assignColorValues(requestedColors models.Menucolors) {
	if requestedColors.MenuForeground != "" {
		MenuFG = render.FgColors[requestedColors.MenuForeground]
	}
	if requestedColors.MenuBackground != "" {
		MenuBG = render.BgColors[requestedColors.MenuBackground]
	}
	if requestedColors.BorderForeground != "" {
		BorderFG = render.FgColors[requestedColors.BorderForeground]
	}
	if requestedColors.BorderBackground != "" {
		BorderBG = render.BgColors[requestedColors.BorderBackground]
	}
}

func assignStyle() {
	MenuStyle = render.New(MenuFG, MenuBG)
	BorderStyle = render.New(BorderFG, BorderBG)
}
