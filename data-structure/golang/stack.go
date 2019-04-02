package dsa

import "errors"

/**
用数组实现一个简单的栈结构，

实现了一些基本的操作：入栈，出栈，获取栈顶元素，判断栈空，栈大小
*/
var StackEmptyError = errors.New("stack is empty")

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(elem interface{}) {
	s.elements = append(s.elements, elem)
}

func (s *Stack) Pop() (interface{}, error) {
	if !s.IsEmpty() {
		pos := s.Size() - 1
		v := s.elements[pos]
		s.elements = s.elements[:pos]
		return v, nil
	}
	return nil, StackEmptyError
}

func (s *Stack) Top() (interface{}, error) {
	if !s.IsEmpty() {
		return s.elements[s.Size()-1], nil
	}
	return nil, StackEmptyError
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}
