// Package meta provides signals and slots for the Observer pattern.
//
// Note: This is a proof of concept package, therefore you are not
// supposed to use it in production or any near. Public API of this
// package is likely to change (a lot).
//
// Signals and slots is a language construct introduced in Qt for
// communication between objects which makes it easy to implement
// the Observer pattern while avoiding boilerplate code. The concept
// is that GUI widgets can send signals containing event information
// which can be received by other controls using special functions,
// known as slots.
//
// Once a signal gets emitted, all the slots connected are being
// executed in the distinct goroutines. Value passed to the signal
// is available from the call context object.
package meta

// Signal is an adorable piece of a structure!
type Signal struct {
	slots []Slot
}

// Slot is a reciver function, usually a wrapper.
type Slot func(*Call)

// Call stands for the call context, handled in the slot.
type Call struct {
	Data interface{}
}

// Connect attaches a new slot to the signal. It also does
// panic if any of the params given is equal to nil.
func Connect(sig *Signal, slot Slot) {
	if sig == nil {
		panic("can't connect a nil signal")
	}

	if slot == nil {
		panic("can't connect a nil slot")
	}

	sig.slots = append(sig.slots, slot)
}

// Emit executes all the connected slots with data given.
func (sig *Signal) Emit(data interface{}) {
	for i := range sig.slots {
		sig.slots[i](&Call{Data: data})
	}
}
