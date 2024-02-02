package anagram

// isAnagram solves exercise 3.10 3-10.
// Does not deal with case thus silent and Listen would not be anagrams.
// Complexity:
// N being the number of characters in a
// M being the number of characters in b
// Time O(N+M)
// Space O(N): for the frequency map. N being bound by the charset/alphabet.
func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	freq := make(map[rune]int)
	for _, r := range a {
		freq[r]++
	}

	for _, r := range b {
		cnt, ok := freq[r]
		if !ok && cnt == 0 {
			return false
		}

		cnt--
		if cnt == 0 {
			delete(freq, r)
		} else {
			freq[r] = cnt
		}
	}

	return len(freq) == 0
}
