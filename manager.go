package xmodule

import (
	//"fmt"
	"log"
	"x-module/hamster"
)

var moduleLst = []*hamster.Hamster{}
var moduleNames = make([]string, 0)

func InitManager() {
	log.Print("...")
}

func NewModule(name string, relyLst []string) {
	m := hamster.NewModule(name, relyLst)
	AddModuleToManage(m)
}

func AddModuleToManage(m *hamster.Hamster) {
	moduleLst = append(moduleLst, m)
	moduleNames = append(moduleNames, hamster.GetModuleName(m))
}

func CheckAllModuleRely() {
	checkLst := make(map[string][]string)
	for _, m := range moduleLst {
		_, unRelyLst := CheckModuleRely(m)
		if len(unRelyLst) != 0 {
			checkLst[hamster.GetModuleName(m)] = unRelyLst
		}
	}
	if len(checkLst) != 0 {
		log.Print("!!!!!!!!!! These are Nil Rely !!!!!!!!!!!!")
		panic("")
	}
}

func CheckModuleRely(m *hamster.Hamster) (relyLst []string, unRelyLst []string) {
	moduleNameLst := GetAllModuleNames()
	relyLst = make([]string, 0)
	unRelyLst = make([]string, 0)
	for _, relyName := range hamster.GetModuleRelyLst(m) {
		var check bool
		for _, moduleName := range moduleNameLst {
			if relyName == moduleName {
				check = true
				break
			}
		}
		if check {
			relyLst = append(relyLst, relyName)
		} else {
			unRelyLst = append(unRelyLst, relyName)
		}
	}
	return relyLst, unRelyLst
}

func GetModule(str string) {

}

func GetAllModuleNames() (nameLst []string) {
	return moduleNames
}

func GetAllModuleRelyLst() (relyLst map[string][]string) {
	relyLst = make(map[string][]string, len(moduleLst))
	for _, m := range moduleLst {
		relyLst[hamster.GetModuleName(m)] = hamster.GetModuleRelyLst(m)
	}
	return relyLst
}
