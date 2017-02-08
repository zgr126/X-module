package hamster

import (
	"github.com/satori/go.uuid"
)

type Hamster struct {
	Hname string
	Huuid string
	Htype string
}

func NewHamster() (h *Hamster) {
	hamster := new(Hamster)
	hamster.Huuid = uuid.NewV1().String()
	return hamster
}

func (h *Hamster) Input(data interface{}) {

}

func (h *Hamster) Output() (data interface{}) {
	return ""
}
