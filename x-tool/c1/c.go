package c1

import (
	"fmt"
	"x-tool/m"
)

var model m.M

func (m *model) control() {
	fmt.Print("c1 !")
}

func init() {
	m.reject(model)
}
