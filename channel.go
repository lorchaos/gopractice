package gopractice

import (
	"log"
)

var tokenChannel = make(chan int, 10)

type Sequencer func() int

func NewSequencer() Sequencer {

	go func() {

		for s := 0; ; {
			s = s + 1
			tokenChannel <- s
		}
	}()

	return func() (t int) {

		t = <-tokenChannel
		log.Printf("Handing token %d", t)
		return
	}
}
