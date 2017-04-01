package main

import (
	_ "x-tool/c1"
	_ "x-tool/c2"
	"x-tool/m"
)

func main() {
	var z = m.get("c2")
	z.control()
}
