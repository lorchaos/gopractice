package gopractice

import (
	"log"
	"testing"
)

func ExampleSequencer() {

	s := NewSequencer()

	if s() == 1 {

		log.Println("This worked!")
	}

}

func TestSequenceGenerator(t *testing.T) {

	s := NewSequencer()

	for i := 1; i < 100000; i++ {

		if r := s(); r != i {

			t.Errorf("Expected %d received %d", i, r)
			t.FailNow()
		}
	}
}
