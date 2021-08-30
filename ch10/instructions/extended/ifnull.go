package extended

import (
	"ch10/instructions/base"
	"ch10/rtda"
)

// IFNULL 判断引用为null
type IFNULL struct {
	base.BranchInstruction
}

// IFNONNULL 判断引用非null
type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
