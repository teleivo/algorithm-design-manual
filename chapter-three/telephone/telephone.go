package telephone

var keypad = map[rune]int{
	'a': 2,
	'b': 2,
	'c': 2,
	'd': 3,
	'e': 3,
	'f': 3,
	'g': 4,
	'h': 4,
	'i': 4,
	'j': 5,
	'k': 5,
	'l': 5,
	'm': 6,
	'n': 6,
	'o': 6,
	'p': 7,
	'q': 7,
	'r': 7,
	's': 7,
	't': 8,
	'u': 8,
	'v': 8,
	'w': 9,
	'x': 9,
	'y': 9,
	'z': 9,
}

// words solves exercise 3.10 3-9.
// Solution runs in O(1) space and O(N*M) time N being the number of digits, M being the number of words in the
// dictionary.
// Assuming all lowercase ASCII characters in the dictionary.
// Assuming that one digit in the input sequence represents one character in the dictionary word.
// Some keypads required pressing a key multiple times to get the correct character like pressing 2
// twice to get 'b'. I am not dealing with such keypads here.
func words(digits []int, dict *list) []string {
	var result []string

	for word := dict.minimum(); word != nil; word = dict.successor(word) {
		if len(word.Value) != len(digits) {
			digit := digits[0]
			char := word.Value[0]
			charDigit := keypad[rune(char)]

			if digit < charDigit {
				// any successor of word cannot be a match
				// for example digit = 2 but the word is 'dog', any successor of 'dog' will at least
				// start with a char equal to 'd' or higher
				return result
			}

			continue
		}
		match := true
		for i, digit := range digits {
			char := word.Value[i]
			charDigit := keypad[rune(char)]

			if digit < charDigit && i == 0 {
				// any successor of word cannot be a match
				// for example digit = 2 but the word is 'dog', any successor of 'dog' will at least
				// start with a char equal to 'd' or higher
				return result
			} else if digit != charDigit {
				match = false
				break
			}
		}
		if match {
			result = append(result, word.Value)
		}
	}

	return result
}

// Copied dictionary implementation from ../list (solves exercise 3.10 3-7)
type list struct {
	Head *node
	Tail *node
}

type node struct {
	Value string
	Next  *node
}

func New(words ...string) *list {
	l := &list{
		Head: &node{},
		Tail: &node{},
	}

	for _, w := range words {
		l.insert(&node{Value: w})
	}
	return l
}

func (l *list) search(x string) *node {
	if l == nil {
		return nil
	}

	for n := l.Head.Next; n != nil; n = n.Next {
		if n.Value == x {
			return n
		}
	}

	return nil
}

func (l *list) insert(x *node) {
	x.Next = l.Head.Next
	l.Head.Next = x
	if x.Next == nil {
		x.Next = l.Tail
	}
}

// delete node x from list in constant time.
func (l *list) delete(x *node) {
	if x.Next == l.Tail {
		l.Tail = x
	}
	x.Value = x.Next.Value
	x.Next = x.Next.Next
}

// predecessor finds the logical predecessor of given node in O(N) time.
func (l *list) predecessor(x *node) *node {
	var result *node

	for n := l.Head.Next; n != nil; n = n.Next {
		if n.Value < x.Value && (result == nil || n.Value > result.Value) {
			result = n
		}
	}

	return result
}

// successor finds the logical successor of given node in O(N) time.
func (l *list) successor(x *node) *node {
	var result *node

	for n := l.Head.Next; n != nil; n = n.Next {
		if n.Value > x.Value && (result == nil || n.Value < result.Value) {
			result = n
		}
	}

	return result
}

// minimum finds the smallest node in O(N) time.
// O(N) as deletion has to be done in O(1) time so I cannot hide the cost of searching in the
// unsorted list in the deletion operation.
func (l *list) minimum() *node {
	var result *node

	for n := l.Head.Next; n != nil && n != l.Tail; n = n.Next {
		if result == nil || n.Value < result.Value {
			result = n
		}
	}

	return result
}

// maximum finds the largest node in O(N) time.
// O(N) as deletion has to be done in O(1) time so I cannot hide the cost of searching in the
// unsorted list in the deletion operation.
func (l *list) maximum() *node {
	var result *node

	for n := l.Head.Next; n != nil; n = n.Next {
		if result == nil || n.Value > result.Value {
			result = n
		}
	}

	return result
}
