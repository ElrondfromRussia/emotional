package emotional

import "log"

const (
	FUCK    = "凸(￣ヘ￣)"
	DISGUST = "(︶︹︺)"
	GRRR    = "٩(ఠ益ఠ)۶"
	GOAWAY  = "╭∩╮(ಠಠ)╭∩╮"
	TAKBLET = "ᕦ(ò_óˇ)ᕤ"
)

func Fuck() {
	log.Println(FUCK)
}

func Disgust() {
	log.Println(DISGUST)
}

func Grrr() {
	log.Println(GRRR)
}

func GoAway() {
	log.Println(GOAWAY)
}

func TakBlet() {
	log.Println(TAKBLET)
}
