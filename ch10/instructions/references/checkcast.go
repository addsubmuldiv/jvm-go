package references

import (
	"ch10/instructions/base"
	"ch10/rtda"
	"ch10/rtda/heap"
)

// Check whether object is of given type
// 这个指令不会改变操作数栈，大概是用于强制转换的判断，能不能转，不能转直接抛出异常
type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
