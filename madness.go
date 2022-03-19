package emotional

import (
	"log"
	"time"
)

const (
	ISTERIC   = "ᕕ| ~ ⁰ ෴ ⁰ ~ |┐"
	APATIC    = "〳 ͡° Ĺ̯ ͡° 〵"
	DEPRESSED = "༼▃ Ĺ̯ ▃༽"
)

func Isteric() {
	log.Fatalln(ISTERIC, RIGHTARROW, GRRR, RIGHTARROW, DISGUST, RIGHTARROW,
		TAKBLET, RIGHTARROW, FUCK, RIGHTARROW, GOAWAY)
}

func Apatic() {
	for true {
		log.Println(APATIC)
	}
}

func Depression() {
	go func() { log.Println(DEPRESSED) }()
	i := 0
	for i < 60 {
		i += 1
		time.Sleep(time.Second)
	}
	Isteric()
}
