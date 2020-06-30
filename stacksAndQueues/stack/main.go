package stack

type stack struct {
	mem []interface{}
}

func MakeStack() *stack {
	return &stack{
		mem: make([]interface{}, 100),
	}
}

func (s *stack) Push(val interface{}) {
	s.mem = append(s.mem, val)
}

func (s *stack) Pop() interface{} {
	ele := s.mem[len(s.mem)-1]
	s.mem = s.mem[:len(s.mem)-1]
	return ele
}

func (s *stack) IsEmpty() bool {
	if len(s.mem) > 0 {
		return false
	}

	return true
}
