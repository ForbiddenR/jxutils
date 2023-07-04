package buffer

import (
	"reflect"
	"testing"
)

func TestGrowth(t *testing.T) {
	t.Parallel()
	x := 10
	g := NewRingGrowing[int](1)
	for i := 0; i < x; i++ {
		if e, a := i, g.readable; !reflect.DeepEqual(e, a) {
			t.Fatalf("expected equal, got %#v, %#v", e, a)
		}
		g.WriteOne(i)
	}
	read := 0
	for g.readable > 0 {
		v, ok := g.ReadOne()
		if !ok {
			t.Fatal("expected true")
		}
		if read != v {
			t.Fatalf("expected %#v==%#v", read, v)
		}
		read++
	}
	if x != read {
		t.Fatalf("expecte to have read %d items: %d", x, read)
	}
	if g.readable != 0 {
		t.Fatalf("expected readable to be zero: %d", g.readable)
	}
	ifg.n  != 16 {
		t.Fatalf("expected N to be 16: %d", g.n)
	}
}