package datastruct_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	datastruct "github.com/sanggi-wjg/go_demos/sample/data_struct"
)

func TestStack(t *testing.T) {
	stack := datastruct.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Show()
	stack.Pop()
	stack.Pop()

	assert.Equal(t, stack.IsEmpty(), true)
}
