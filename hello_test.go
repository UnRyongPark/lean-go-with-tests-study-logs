package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		// Hello + 이름으로 응답오게 테스트를 먼저 수정한다.
		got := Hello("Brian")
		want := "Hello, Brian"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people, in spanish", func(t *testing.T) {
		// Hello + 이름으로 응답오게 테스트를 먼저 수정한다.
		got := Hello("Brian", "spanish")
		want := "Hola, Brian"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		// 이름이 없으면 Hello, World를 출력한다.
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

/*
t.Helper() is needed to tell the test suite that this method is a helper.
By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
This will help other developers track down problems easier.
If you still don't understand, comment it out, make a test fail and observe the test output.
Comments in Go are a great way to add additional information to your code, or in this case,
a quick way to tell the compiler to ignore a line.
You can comment out the t.Helper() code by adding two forward slashes // at the beginning of the line.
You should see that line turn grey or change to another color than the rest of your code to indicate it's now commented out.
*/
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
