package main

import "testing"

func TestHello(t *testing.T) {
	// Hello + 이름으로 응답오게 테스트를 먼저 수정한다.
	got := Hello("Brian")
	want := "Hello, Brian"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
