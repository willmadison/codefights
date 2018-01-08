package codefights

import (
	"bytes"
	"errors"
)

type Stack interface {
	Push(b *bytes.Buffer)
	Pop() (*bytes.Buffer, error)
	Peek() (*bytes.Buffer, error)
	IsEmpty() bool
}

type bufferStack struct {
	data []*bytes.Buffer
	size int
}

func (s *bufferStack) Push(b *bytes.Buffer) {
	s.data = append(s.data, b)
	s.size++
}

func (s *bufferStack) Pop() (*bytes.Buffer, error) {
	if s.size > 0 {
		value := s.data[s.size-1]
		s.size--
		s.data = s.data[:s.size]
		return value, nil
	}

	return nil, errors.New("no such element")
}

func (s *bufferStack) Peek() (*bytes.Buffer, error) {
	if s.size > 0 {
		return s.data[s.size-1], nil
	}

	return nil, errors.New("no such element")
}

func (s *bufferStack) IsEmpty() bool {
	return s.size == 0
}

func NewBufferStack() Stack {
	return &bufferStack{data: []*bytes.Buffer{}}
}

func reverseParenthesis(s string) string {
	var buffer bytes.Buffer

	stack := NewBufferStack()

	for _, character := range s {
		switch character {
		case '(':
			stack.Push(bytes.NewBuffer(nil))
		case ')':
			b, _ := stack.Pop()
			reversed := reverse(b.String())

			if !stack.IsEmpty() {
				b, _ = stack.Peek()
				b.WriteString(reversed)
			} else {
				buffer.WriteString(reversed)
			}
		default:
			if stack.IsEmpty() {
				buffer.WriteRune(character)
			} else {
				b, err := stack.Peek()
				if err != nil {
					panic(err)
				}
				b.WriteRune(character)
			}
		}
	}

	return buffer.String()
}

func reverse(s string) string {
	var b bytes.Buffer

	for i := len(s) - 1; i >= 0; i-- {
		b.WriteRune(rune(s[i]))
	}

	return b.String()
}
