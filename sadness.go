package emotional

import (
	"log"
)

const (
	CRY      = "(ಥ﹏ಥ)"
	SAD      = "(ಥ_ಥ)"
	DEAD     = "(✖╭╮✖)"
	DEADLINE = "(☆_@)"
)

func Cry() {
	log.Println(CRY)
}

func Sad() {
	log.Println(SAD)
}

func Dead() {
	log.Println(DEAD)
}

func DeadLine() {
	log.Println(DEADLINE)
}
