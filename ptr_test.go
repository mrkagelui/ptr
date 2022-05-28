package ptr

import (
	"testing"
)

func equals[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestOf(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		num := 24
		equals(t, *Of(num), num)
	})
	t.Run("int8", func(t *testing.T) {
		var num int8 = 24
		equals(t, *Of(num), num)
	})
	t.Run("int16", func(t *testing.T) {
		var num int16 = 24
		equals(t, *Of(num), num)
	})
	t.Run("int32", func(t *testing.T) {
		var num int32 = 24
		equals(t, *Of(num), num)
	})
	t.Run("int64", func(t *testing.T) {
		var num int64 = 24
		equals(t, *Of(num), num)
	})
	t.Run("uint", func(t *testing.T) {
		var num uint = 24
		equals(t, *Of(num), num)
	})
	t.Run("uint8", func(t *testing.T) {
		var num uint8 = 24
		equals(t, *Of(num), num)
	})
	t.Run("uint16", func(t *testing.T) {
		var num uint16 = 24
		equals(t, *Of(num), num)
	})
	t.Run("uint32", func(t *testing.T) {
		var num uint32 = 24
		equals(t, *Of(num), num)
	})
	t.Run("uint64", func(t *testing.T) {
		var num uint64 = 24
		equals(t, *Of(num), num)
	})
	t.Run("uintptr", func(t *testing.T) {
		var num uintptr = 24
		equals(t, *Of(num), num)
	})
	t.Run("float32", func(t *testing.T) {
		var num float32 = 24
		equals(t, *Of(num), num)
	})
	t.Run("float64", func(t *testing.T) {
		var num float64 = 24
		equals(t, *Of(num), num)
	})
	t.Run("complex64", func(t *testing.T) {
		var num complex64 = 24 + 3i
		equals(t, *Of(num), num)
	})
	t.Run("complex128", func(t *testing.T) {
		num := 24.42 - 4i
		equals(t, *Of(num), num)
	})
	t.Run("byte", func(t *testing.T) {
		var b byte = 'a'
		equals(t, *Of(b), b)
	})
	t.Run("rune", func(t *testing.T) {
		var r rune = '义'
		equals(t, *Of(r), r)
	})
	t.Run("string", func(t *testing.T) {
		title := "maidenless"
		equals(t, *Of(title), title)
	})
	t.Run("boolean", func(t *testing.T) {
		b := true
		equals(t, *Of(b), b)
	})
	t.Run("array", func(t *testing.T) {
		var arr = [9]int{32, 921}
		equals(t, *Of(arr), arr)
	})
	t.Run("struct", func(t *testing.T) {
		type player struct {
			title string
		}
		p := player{"foul tarnished"}
		equals(t, *Of(p), p)
	})
	t.Run("channel", func(t *testing.T) {
		m := make(chan bool)
		equals(t, *Of(m), m)
	})
}
