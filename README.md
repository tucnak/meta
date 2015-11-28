>This is a proof of concept implementation of signals and slots in Go.

[![GoDoc](https://godoc.org/github.com/tucnak/meta?status.svg)](https://godoc.org/github.com/tucnak/meta)

Signals and slots is a language construct introduced in Qt for communication between objects, which makes it easy to implement the Observer pattern while avoiding boilerplate code. The concept is that GUI widgets can send signals containing event information which can be received by other controls using special functions known as slots. Similarly, the signal/slot system can be used for other non-GUI usages, for example asynchronous I/O (including sockets, pipes, serial devices, etc.) event notification or to associate timeout events with appropriate object instances and methods or functions.

A commonly used metaphor is a spreadsheet. A spreadsheet has cells that observe the source cell(s). When the source cell is changed, the dependent cells are updated from the event. Channels can't do that: you can read a value from the channel at some particular place only. As you can see, it's not flexible enough, especially for the systems with lots of observers — e.g. user interfaces.

Since Go does not support generics and we don't have any sort of a Meta Object Compiler from Qt, there is absolutely no way to omit the boilerplate code:

```go
package main

import (
	"fmt"
	"github.com/tucnak/meta"
	"sync"
)

type Foo struct {
	// fields...

	Done meta.Signal
}

type Guy struct {
	Name string
}

func (mr Guy) Print(n int) {
	fmt.Printf("%s says: %d\n", mr.Name, n)
	waiter.Done()
}

var waiter sync.WaitGroup

func main() {
	var foo Foo

	johny := Guy{"Johny"}
	david := Guy{"David"}

	meta.Connect(&foo.Done, func(call *meta.Call) {
		if passed, ok := call.Data.(int); ok {
			johny.Print(passed)
		}
	})

	meta.Connect(&foo.Done, func(call *meta.Call) {
		if passed, ok := call.Data.(int); ok {
			david.Print(passed)
		}
	})

	waiter.Add(2)

	// Emit notifies all the connected slots, by running
	// them in the distinct goroutines.
	foo.Done.Emit(42)

	waiter.Wait()

	//
	// Johny says: 42
	// David says: 42
}
```

Here it is, feel free to contribute.
