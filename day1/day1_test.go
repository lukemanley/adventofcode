package main

import "testing"

func TestP1(t *testing.T) {
	floor := p1()
	if floor != 232 {
		t.Errorf("expected %d, got %d", 232, floor)
	}
}

func TestP2(t *testing.T) {
	floor := p2()
	if floor != 1783 {
		t.Errorf("expected %d, got %d", 1783, floor)
	}
}
