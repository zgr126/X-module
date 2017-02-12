package xmodule

import (
	"x-module/hamster/configRead"
	"x-module/hamster/httpListen"
	"x-module/hamster/router"
)

func InitBaseHamster() {
	_ = router.init()
	_ = httpListen.init()
	_ = configRead.init()
}
