package palindrome

// IsPalindrome accepts a string and returns true if it is a palindrome
func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] == s[len(s)-1] {
		return IsPalindrome(s[1 : len(s)-1])
	}

	return false
}
