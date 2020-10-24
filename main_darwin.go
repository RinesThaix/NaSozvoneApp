package main

import (
	"github.com/caseymrm/menuet"
)

func main() {
	go addToMenuBar()
	go Run()
	menuet.App().RunApplication()
}

func addToMenuBar() {
	menuet.App().SetMenuState(&menuet.MenuState{
		Image: "icon_512x512.png",
	})
}
