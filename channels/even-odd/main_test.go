package numbernoise

import (
	"testing"
	"time"
)

func TestRandomNSeconds(t *testing.T) {
	t.Run("Test that we get a stream of numbers > 0", func(t *testing.T) {
		nums := RandomNSeconds(time.Millisecond, 50*time.Millisecond)
		ok := true

		for ok {
			n, ok := <-nums
			if !ok {
				break
			}

			if n == 0 {
				t.Logf("Zero value received from random number generator while the channel was open: %d\n", n)
				t.Fail()
			}
		}
	})
}

func TestEvenOdds(t *testing.T) {
	evens, odds := EvenOdds(time.Millisecond, 2*time.Second)
	ti := time.NewTimer(1 * time.Second)
	e := make([]int, 2000)
	o := make([]int, 2000)

	ok := true

	for ok {
		select {
		case en := <-evens:
			e = append(e, en)
		case on := <-odds:
			o = append(o, on)
		case <-ti.C: // this emits after the time limit
			ok = false
		}
	}

	t.Run("Only even numbers in evens", func(t *testing.T) {
		for _, i := range e {
			if i == 0 {
				break // channel is closed
			}

			if i%2 != 0 {
				t.Logf("Uneven number in evens: %d", i)
				t.Fail()
			}
		}
	})

	t.Run("Only odd numbers in odds", func(t *testing.T) {
		for _, i := range o {
			if i == 0 {
				break // channel is closed
			}

			if i%2 > 0 || i%2 < 0 {
			} else {
				t.Logf("Even number in odds: %d", i)
				t.Fail()
			}
		}
	})

}
