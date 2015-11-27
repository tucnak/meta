package meta

import (
	"testing"
	"time"
)

type A struct {
	done Signal
}

type B struct{}

func (B) sum(a, b int) int {
	return a + b
}

type twoNumbers struct {
	a, b int
}

func TestConnect(t *testing.T) {
	var a A
	var b B

	pipe := make(chan int)

	Connect(&a.done, func(call *Call) {
		if nums, ok := call.Data.(twoNumbers); ok {
			pipe <- b.sum(nums.a, nums.b)
		} else {
			t.Fatal("transferred data is rubbish")
			return
		}
	})

	go func() {
		<-time.After(10 * time.Millisecond)
		a.done.Emit(twoNumbers{40, 2})
	}()

	x := <-pipe

	if x != 42 {
		t.Fatal("the whole thing doesn't really work")
		return
	}
}
