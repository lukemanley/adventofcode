package main

import "testing"

func TestP1(t *testing.T) {
	expected := 1448828
	got := p1()
	if expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}
}

func TestP2(t *testing.T) {
	expected := 3783758
	got := p2()
	if expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}
}
