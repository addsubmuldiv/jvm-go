package loads

// int类型

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

// ILOAD Load int from local variable
// 注意这里用的嵌套是 base.Index8Instruction，说明它要读取一个字节的操作数
// 下面几个方法则是 索引隐含在指令中了
type ILOAD struct{ base.Index8Instruction }

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, self.Index)
}

type ILOAD_0 struct{ base.NoOperandsInstruction }

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperandsInstruction }

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
