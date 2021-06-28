package comparisons

import (
	"ch07/instructions/base"
	"ch07/rtda"
)

// DCMPG Compare double
type DCMPG struct{ base.NoOperandsInstruction }

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

// DCMPL 和float差不多, 整俩为了NaN
type DCMPL struct{ base.NoOperandsInstruction }

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
