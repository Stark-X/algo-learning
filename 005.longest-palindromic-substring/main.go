package main

func longestPalindrome(s string) string {
	leftBorder, rightBorder, maxLength := 0, 0, 0
	left, right, length := 0, 0, 1
	for i := 0; i < len(s); i++ {
		left = i - 1
		right = i + 1
		for left >= 0 && s[left] == s[i] {
			left--
			length++
		}
		for right < len(s) && s[right] == s[i] {
			right++
			length++
		}
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			length += 2
		}
		if length > maxLength {
			leftBorder = left
			rightBorder = right
			maxLength = length
		}
		length = 1
	}

	return s[leftBorder+1 : rightBorder]
}
