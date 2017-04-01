package c2

import (
	"fmt"
	"x-tool/m"
)

var model m.M

func (mo *model) control() {
	var z = m.get("c1")
	z.control()
	fmt.Print("c2 !")
}

func init() {
	m.reject(model)
}
