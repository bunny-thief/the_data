package palindrome

// IsPalindrome accepts a string and returns true if it is a palindrome
func IsPalindrome(s string) bool {
	if s[0] == s[len(s)-1] {
		if len(s) <= 2 {
			return true
		}
		return IsPalindrome(s[1 : len(s)-1])
	}
	return false
}
