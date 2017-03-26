// struct module to manger
package xmodule

import (
	"fmt"
)

type ModuleManage struct {
	Hname    string
	Huuid    string
	Htype    string
	HrelyLst []string
	hamster.Hamster
}

var moduleLst = []*ModuleManage{}
var moduleNames = make([]string, 0)

func InitManager() {
	InitBaseHamster()
	log.Print("...")
}

func NewModule(name string, relyLst []string) {
	m := NewHamster(name, relyLst)
	AddModuleToManage(m)
}

func AddModuleToManage(m *ModuleManage) {
	moduleLst = append(moduleLst, m)
	moduleNames = append(moduleNames, GetModuleName(m))
}

func NewHamster(name string, relyLst []string) (h *ModuleManage) {
	hamster := new(ModuleManage)
	hamster.Huuid = uuid.NewV1().String()
	hamster.Hname = name
	hamster.HrelyLst = relyLst
	// if hamster has no rely, type is base hamster
	if len(relyLst) == 0 {
		hamster.Htype = "base"
	}
	return hamster
}

func GetModuleName(m *ModuleManage) (Hname string) {
	return m.Hname
}

func GetModuleRelyLst(m *ModuleManage) (HrelyLst []string) {
	return m.HrelyLst
}

func CheckAllModuleRely() {
	checkLst := make(map[string][]string)
	for _, m := range moduleLst {
		_, unRelyLst := CheckModuleRely(m)
		if len(unRelyLst) != 0 {
			checkLst[GetModuleName(m)] = unRelyLst
		}
	}
	if len(checkLst) != 0 {
		log.Print("!!!!!!!!!! These are Nil Rely !!!!!!!!!!!!")
		panic("")
	}
}

func CheckModuleRely(m *ModuleManage) (relyLst []string, unRelyLst []string) {
	moduleNameLst := GetAllModuleNames()
	relyLst = make([]string, 0)
	unRelyLst = make([]string, 0)
	for _, relyName := range GetModuleRelyLst(m) {
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

func GetModuleUseName(str string) {

}

func GetAllModuleNames() (nameLst []string) {
	return moduleNames
}

func GetAllModuleRelyLst() (relyLst map[string][]string) {
	relyLst = make(map[string][]string, len(moduleLst))
	for _, m := range moduleLst {
		relyLst[GetModuleName(m)] = GetModuleRelyLst(m)
	}
	return relyLst
}
