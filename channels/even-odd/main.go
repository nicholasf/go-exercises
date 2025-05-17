// An exercise that takes ever n seconds from Time.tick then

// do time.Tick millisecond then for each result multiply it by a random seed, you'll ha a unique result, then put them into a channel,
// then see if you can select on them by even and odd.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedMillsecond = int

func main() {
	c := time.Tick(time.Millisecond)

	evens := make(chan RandomizedMillsecond, 100000)
	odds := make(chan RandomizedMillsecond, 100000)

	go func() {
		for tick := range c {
			t := int(int64(tick.UnixMilli())) * rand.Int()

			if t%2 == 0 {
				evens <- t
			} else {
				odds <- t
			}
		}
	}()

	var e, o = 0, 0

	for {
		select {
		case n := <-evens:
			{
				e = e + 1
				fmt.Println("EVEN: ", n, " Count: ", e)
			}
		case n := <-odds:
			{
				o = o + 1
				fmt.Println("ODD: ", n, " Count: ", o)
			}
		}
	}
}
