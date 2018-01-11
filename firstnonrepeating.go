package codefights

import "errors"

type RuneStack interface {
	Push(r rune)
	Pop() (rune, error)
	Peek() (rune, error)
	IsEmpty() bool
}

type runeStack struct {
	data []rune
	size int
}

func NewRuneStack() RuneStack {
	return &runeStack{data: make([]rune, 26)}
}

func (s *runeStack) Push(value rune) {
	s.data[s.size] = value
	s.size++
}

func (s *runeStack) Pop() (rune, error) {
	if s.size > 0 {
		var empty rune
		value := s.data[s.size-1]
		s.data[s.size-1] = empty
		s.size--
		return value, nil
	}

	return 0, errors.New("No Such Element")
}

func (s *runeStack) Peek() (rune, error) {
	if s.size > 0 {
		value := s.data[s.size-1]
		return value, nil
	}

	return 0, errors.New("No Such Element")
}

func (s *runeStack) IsEmpty() bool {
	return s.size == 0
}

func firstNotRepeatingCharacter(s string) string {
	var bitMask int

	stack := NewRuneStack()

	lastUniqueSeen, previousLastUnique := '_', '_'

	for index := len(s) - 1; index >= 0; index-- {
		c := rune(s[index])
		if seen := isBitSet(bitMask, int(c-'a')); !seen {
			lastUniqueSeen, previousLastUnique = c, lastUniqueSeen
			setBit(&bitMask, int(c-'a'))
			stack.Push(previousLastUnique)
		} else if lastUniqueSeen == c {
			lastUniqueSeen, _ = stack.Pop()
		}
	}

	return string(lastUniqueSeen)
}

func setBit(given *int, i int) {
	*given |= (1 << uint(i))
}

func isBitSet(given int, i int) bool {
	return given&(1<<uint(i)) != 0
}
