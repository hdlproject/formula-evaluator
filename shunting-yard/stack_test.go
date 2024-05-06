package shunting_yard

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStringStack(t *testing.T) {
	input := []string{"1", "2", "3"}
	expectedOutput := []string{"3", "2", "1"}

	stack := stringStack{}
	for _, item := range input {
		stack.Push(item)
	}

	var output []string
	item, isFound := stack.Pop()
	for isFound {
		output = append(output, item)

		item, isFound = stack.Pop()
	}

	if diff := cmp.Diff(expectedOutput, output); diff != "" {
		t.Fatalf("(-want/+got)/n%s", diff)
	}
}
