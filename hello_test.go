package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		// Hello + 이름으로 응답오게 테스트를 먼저 수정한다.
		got := Hello("Brian")
		want := "Hello, Brian"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		// 이름이 없으면 Hello, World를 출력한다.
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
