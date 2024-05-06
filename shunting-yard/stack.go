package shunting_yard

type stringStack []string

func (s *stringStack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stringStack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}

	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item, true
}

func (s *stringStack) Push(item string) {
	*s = append(*s, item)
}
