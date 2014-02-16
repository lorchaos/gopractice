package gopractice

import ("testing"
		"fmt"
)

func ExampleSequencer() {

	s := NewSequencer()

	end := make(chan int, 10)

	f := func(i int) {

		for i := 0; i < 25000; i++ {

			s();
		}
		end <- i
	}

	//let's call the sequencer from different goroutines
	for i := 0; i < 4; i++ {

		go f(i)
	}

	//so the test won't exit 
	for i := 0; i < 4; i++ {
		end <- i
		fmt.Printf("Goroutine %d ended\n", i)
	}

	//it's interesting how the goroutines always finish in the order started
	//Output: Goroutine 0 ended
	//Goroutine 1 ended
	//Goroutine 2 ended
	//Goroutine 3 ended
}


func TestSequenceGenerator(t *testing.T) {

	s := NewSequencer()

	for i := 0; i < 100000; i++ {
			
		if r := s(); r != i {

			t.Errorf("Expected %d received %d", i, r)
			t.FailNow()
		}
	}
}
