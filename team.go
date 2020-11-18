package emotional

import "log"

const (
	BACKEND  = "ヽ༼ ຈ̚ل͜ʘ̚༼Ɵ͆_◕༽◉_◔ ༽ﾉ"
	NEWGUYS  = "༼ ºل͟º༼ ºل͟º( ͡° ͜ʖ ͡°)ºل͟º ༽ºل͟º ༽"
	FRONTEND = "ヽ༼ Ɵ͆ل͜Ɵ͆̚༼◕_Ɵ͆༽◕_◔ ༽ﾉ"
)

func Backend() {
	log.Println(BACKEND)
}

func NewGuys() {
	log.Println(NEWGUYS)
}

func Frontend() {
	log.Println(FRONTEND)
}
