package main

import (
	"GoFormatMenu/menu"
	"GoFormatMenu/models"
	"strings"
)

/*
TODO
	Interactive Mode
	Change Colors
	JSON Validation
*/

var (
	mainMenu = `Menu_Files\Mainmenu.json`
)

func main() {
	currentMenu := models.GetMenu(&mainMenu)

	for {
		index := menu.WriteMenu(&currentMenu)

		if currentMenu.MenuItems[index].Key == `Exit` {
			break
		}

		if strings.HasSuffix(currentMenu.MenuItems[index].Action, `.json`) {
			currentMenu = models.GetMenu(&currentMenu.MenuItems[index].Action)
		} else {
			menu.Execute(&currentMenu.MenuItems[index].Action)
		}
	}
}
