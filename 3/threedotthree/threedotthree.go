package threedotthree

import (
	"github.com/Analyse4/mew/adt/stack"
)

// isStackFull should implement in package stack, but package mew/stack is variable length stack
// add fixedSize check stack if is full
const fixedSize = 5

type SetOfStacks struct {
	stacks []*stack.Stack
}

func New() *SetOfStacks {
	return &SetOfStacks{make([]*stack.Stack, 0)}
}

func (ss *SetOfStacks) Push(v interface{}) {
	if ss.IsEmpty() || ss.stacks[len(ss.stacks)-1].Size() >= fixedSize {
		ss.stacks = append(ss.stacks, new(stack.Stack))
	}
	ss.stacks[len(ss.stacks)-1].Push(v)
}

func (ss *SetOfStacks) Pop() interface{} {
	v := ss.stacks[len(ss.stacks)-1].Pop()
	if ss.stacks[len(ss.stacks)-1].Size() == 0 {
		ss.stacks = ss.stacks[:len(ss.stacks)-1]
	}
	return v
}

func (ss *SetOfStacks) IsEmpty() bool {
	return len(ss.stacks) == 0
}

func (ss *SetOfStacks) Size() int {
	return len(ss.stacks)
}
