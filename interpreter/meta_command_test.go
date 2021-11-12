package interpreter

import "testing"

func TestIsMetaCommand(t *testing.T) {
	assertResult := func(t testing.TB, input string, want bool) {
		t.Helper()

		got := isMetaCommand(input)

		if got != want {
			t.Errorf("isMetaCommand(\"%v\"). got: %t, want: %t", input, got, want)
		}
	}

	t.Run("empty string", func(t *testing.T) {
		t.Parallel()
		assertResult(t, "", false)
	})

	t.Run("the string: \".\"", func(t *testing.T) {
		t.Parallel()
		assertResult(t, ".", true)
	})

	t.Run("a string that does start with '.'", func(t *testing.T) {
		t.Parallel()
		assertResult(t, ".test", true)
	})

	t.Run("a string that does not start with '.'", func(t *testing.T) {
		t.Parallel()
		assertResult(t, "test", false)
	})
}
