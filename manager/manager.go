package manager

import (
	//"fmt"
	"log"
	"x-module/hamster"
)

var moduleLst = []*hamster.Hamster{}

func InitManager() {
	log.Print("...")
}

func NewModule(name string, relyLst []string) {
	m := hamster.NewModule(name, relyLst)
	moduleLst = append(moduleLst, m)
}
func CheckRely(m *hamster.Hamster) bool {
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
	nameLst = make([]string, len(moduleLst))
	for i, m := range moduleLst {
		//fmt.Print(nameLst)
		nameLst[i] = m.Hname
	}
	return nameLst
}
