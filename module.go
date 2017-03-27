package xmodule

import (
	"log"
	"time"
)

type module struct {
	particle
}

// inject instance function struct
type particle struct {
	pTime time.Time
}

// Insert particle
func (p *particle) Insert () (particle) {
	p.pTime = time.Now()
	return p
}

// remove particle
func (p *particle) Remove (pr particle)  {
	log.Print(pr.pTime.Nanosecond())
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
