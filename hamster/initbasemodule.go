package hamster

import (
	// "x-module/hamster/configRead"
	// "x-module/hamster/httpListen"
	"x-module/hamster/router"
)

func init() {
	_ = router.init()
	// _ = httpListen
	// _ = configRead
}
