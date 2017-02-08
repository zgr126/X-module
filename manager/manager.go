package manager

import (
	"log"
	"x-module/hamster"
)

type module hamster.Hamster

var moduleLst []module

func InitManager() {
	log.Print("...")
}

func NewModule(name string, relyLst []string) {
	relyLst()
	m := hamster.NewModule(name, relyLst)
}
func CheckRely(m *module) {
	nameLst := GetModuleNames()
	for _, name := range nameLst {
		if m.Hname == name {
			return true
		}
	}
	panic("!!!!!!! ERROR :" + m.Hname + "HAS NO RELY !!!!!!!!!!")
}

func GetModule(str string) {

}

func GetModuleNames() (nameLst []string) {
	nameLst = nameLst[:len(moduleLst)]
	for i, m := range moduleLst {
		name[i] = m.Hname
	}
	return nameLst
}
