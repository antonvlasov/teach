package main

import (
	"testing"
)

type Stringer interface {
	String() string
}

func GetString(s Stringer) string {
	return s.String()
}

type Text struct {
	val string
	pad [24]byte
}

func (t Text) String() string {
	return t.val
}

func BenchmarkGetString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := Text{val: "hello"} // Moved to heap, 24 bytes for pad and 8 bytes for string.
		_ = GetString(&t)       // Allocation for Stringer 8 + 8 bytes.
	}
}
