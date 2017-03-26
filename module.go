package xmodule

import (
	"log"
)

type module struct {
	//xmodule.ModuleManage
}

// new module struct
func newModule()(module)  {
	newModule := new module
	return newModule
}

func (h *module) Input(data interface{}) {

}

func (h *module) Output() (data interface{}) {
	return ""
}
