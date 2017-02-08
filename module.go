package xmodule

import (
	"fmt"
	"x-module/manager"
)

func Start() {
	manager.InitManager()
	fmt.Print("Strat Success")
}
