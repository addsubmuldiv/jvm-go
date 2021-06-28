package references

import (
	"ch07/instructions/base"
	"ch07/rtda"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
