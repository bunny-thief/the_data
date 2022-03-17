package palindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, expected bool, testString string) {
		t.Helper()
		if got != expected {
			t.Errorf("got %v expected %v, given %q", got, expected, testString)
		}
	}

	t.Run("'kayak' - true", func(t *testing.T) {
		testString := "kayak"
		got := IsPalindrome(testString)
		expected := true
		assertCorrectMessage(t, got, expected, testString)
	})

	t.Run("'car' - false", func(t *testing.T) {
		testString := "car"
		got := IsPalindrome("car")
		expected := false
		assertCorrectMessage(t, got, expected, testString)
	})
}

func ExampleIsPalindrome() {
	palindrome := IsPalindrome("rotor")
	fmt.Println(palindrome)
	// Output: true
}
