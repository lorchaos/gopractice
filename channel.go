package gopractice

import ("log"
)

type sequencer func() int

var _ = log.Printf


func NewSequencer() sequencer {

	tokenChannel := make(chan int, 10)

	go func() {

		for s := 0; ; s++ {
			tokenChannel <- s
		}
	}()

	return func() (int) {

		return <-tokenChannel
	}
}
