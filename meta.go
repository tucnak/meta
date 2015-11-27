package meta

// Signal ...
type Signal struct {
	parent interface{}
	slots  []Slot
}

// Slot ...
type Slot func(*Call)

// Call ...
type Call struct {
	Data interface{}
}

// Connect ...
func Connect(sig *Signal, slot Slot) {
	if sig == nil {
		panic("can't connect a nil signal")
	}

	if slot == nil {
		panic("can't connect a nil slot")
	}

	sig.slots = append(sig.slots, slot)
}

// Emit ...
func (sig *Signal) Emit(data interface{}) {
	for i := range sig.slots {
		sig.slots[i](&Call{Data: data})
	}
}
