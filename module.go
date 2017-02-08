package xmodule

import (
	"fmt"
	"x-module/hamster"
	"x-module/manager"
)

func Start() {
	manager.InitManager()
	fmt.Print("Start Success")
}

func NewModule(name string, relyLst []string) {
	manager.NewModule(name, relyLst)
}
