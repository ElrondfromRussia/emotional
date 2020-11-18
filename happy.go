package emotional

import (
	"log"
)

const (
	HAPPY    = "(＾▽＾)"
	HAPPYCAT = "(─‿‿─)"
	LOVE     = "(─‿‿─)♡"
)

func Happy() {
	log.Println(HAPPY)
}

func HappyCat() {
	log.Println(HAPPYCAT)
}

func Love() {
	log.Println(LOVE)
}
