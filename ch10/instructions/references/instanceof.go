package references

import (
	"ch10/instructions/base"
	"ch10/rtda"
	"ch10/rtda/heap"
)

type INSTANCE_OF struct {
	base.Index16Instruction
}

// 两个操作数，一个是类符号引用，一个是从操作数栈里面弹出的对象
func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass() // 根据符号引用解析出类
	if ref.IsInstanceOf(class) {
		stack.PushInt(1) // 判断结果压栈
	} else {
		stack.PushInt(0)
	}
}
