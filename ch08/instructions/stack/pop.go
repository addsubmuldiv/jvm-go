package stack

import (
	"ch08/instructions/base"
	"ch08/rtda"
)

// POP 用于弹出占用一个操作数栈位置的int、float
type POP struct {
	base.NoOperandsInstruction
}

// POP2 用于弹出占俩操作数栈位置的double、long
type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
