package main

import "testing"

func TestTheMain(t *testing.T) {
	name := getUserName()
	if name == "" {
		t.Errorf("expected username but got nothing!")
	}
}
