package emotional

import (
	elogging "github.com/ElrondfromRussia/emotional/logging"
	"log"
	"time"
)

func Isteric() {
	go func() { elogging.SirLogger.Fatalln("ᕕ| ~ ⁰ ෴ ⁰ ~ |┐") }()
	go func() { elogging.SirLogger.Fatalln("ᕕ| ~ ⁰ ෴ ⁰ ~ |┐") }()
	go func() { log.Fatalln("ᕕ| ~ ⁰ ෴ ⁰ ~ |┐") }()
}

func Apatic() {
	for true {
		log.Println("〳 ͡° Ĺ̯ ͡° 〵")
	}
}

func Depression() {
	go func() { log.Println("༼▃ Ĺ̯ ▃༽") }()
	i := 0
	for i < 60 {
		i += 1
		time.Sleep(time.Second)
	}
	Isteric()
}
