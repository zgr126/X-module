package hamster

import (
	"github.com/satori/go.uuid"
)

type Hamster struct {
	Hname    string
	Huuid    string
	Htype    string
	HrelyLst []string
}

func (h *Hamster) Input(data interface{}) {

}

func (h *Hamster) Output() (data interface{}) {
	return ""
}

func NewModule(name string, relyLst []string) (h *Hamster) {
	hamster := new(Hamster)
	hamster.Huuid = uuid.NewV1().String()
	hamster.Hname = name
	hamster.HrelyLst = relyLst
	// if hamster has no rely, type is base hamster
	if len(relyLst) {
		hamster.Htype = "base"
	}
	return hamster
}
