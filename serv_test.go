package main

import "testing"

func TestServ(t *testing.T) {
	got := greeting("Code.Education Rocks!")
	if got != "<b>Code.Education Rocks!</b>" {
		t.Errorf("got: %s, want: %s", got, "<b>Code.Education Rocks!</b>")
	}
}