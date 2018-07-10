package stack

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
