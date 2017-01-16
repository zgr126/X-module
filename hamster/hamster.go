package hamster

import (
	"github.com/satori/go.uuid"
	"x-module/manager"
)

type Hamster struct {
	Name string
	uuid string
}

func NewHamster() {
	hamster = new(Hamster)
	Hamster.uuid = uuid.NewV1().String()
	return hamster
}
