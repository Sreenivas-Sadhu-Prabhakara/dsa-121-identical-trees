package main

import "testing"

func TestIdentical(t *testing.T) {
	a := tn(1, tn(2, nil, nil), tn(3, nil, nil))
	b := tn(1, tn(2, nil, nil), tn(3, nil, nil))
	if !Identical(a, b) {
		t.Error("equal trees reported different")
	}
	// Same values, different shape.
	c := tn(1, tn(2, nil, nil), nil)
	d := tn(1, nil, tn(2, nil, nil))
	if Identical(c, d) {
		t.Error("different shapes reported identical")
	}
	// Same shape, different value.
	e := tn(1, tn(9, nil, nil), tn(3, nil, nil))
	if Identical(a, e) {
		t.Error("different values reported identical")
	}
	if !Identical(nil, nil) {
		t.Error("two nils are identical")
	}
	if Identical(a, nil) {
		t.Error("tree vs nil reported identical")
	}
}
