package brackets

type (
	Stack struct {
		head *node
		len  int
	}
	node struct {
		val  interface{}
		prev *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Push(val interface{}) {
	n := &node{val, s.head}
	s.head = n
	s.len++
}

func (s *Stack) Pop() interface{} {
	if s.len == 0 {
		return nil
	}
	n := s.head
	s.head = n.prev
	s.len--
	return n.val
}

func (s *Stack) Len() int {
	return s.len
}

func Bracket(str string) (bool, error) {
	var stack *Stack = New()
	for _, v := range str {
		if v == '(' {
			stack.Push(')')
		} else if v == '{' {
			stack.Push('}')
		} else if v == '[' {
			stack.Push(']')
		} else {
			item := stack.Pop()
			if item == nil || item != v {
				return false, nil
			}
		}
	}
	return stack.Len() == 0, nil
}
