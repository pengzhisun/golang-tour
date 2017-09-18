/* ****************************************************************************
 * Tour: sync.Mutex
 * Link: https://tour.golang.org/concurrency/9
 * ----------------------------------------------------------------------------
 * We've seen how channels are great for communication among goroutines.
 *
 * But what if we don't need communication? What if we just want to make sure
 * only one goroutine can access a variable at a time to avoid conflicts?
 *
 * This concept is called mutual exclusion, and the conventional name for the
 * data structure that provides it is mutex.
 *
 * Go's standard library provides mutual exclusion with sync.Mutex and its two
 * methods:
 *   Lock
 *   Unlock
 *
 * We can define a block of code to be executed in mutual exclusion by
 * surrounding it with a call to Lock and Unlock as shown on the Inc method.
 *
 * We can also use defer to ensure the mutex will be unlocked as in the Value
 * method.
 * ***************************************************************************/

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter struct
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc method for SafeCounter struct
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()

	c.v[key]++
	c.mux.Unlock()
}

// Value method for SafeCounter struct
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()

	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
