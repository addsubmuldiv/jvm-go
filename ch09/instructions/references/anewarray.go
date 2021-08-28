package references

import (
	"ch09/instructions/base"
	"ch09/rtda"
	"ch09/rtda/heap"
)

// Create new array of reference
type ANEW_ARRAY struct{ base.Index16Instruction } // 这个操作数是类符号引用，即数组元素的类

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()

	// if componentClass.InitializationNotStarted() {
	// 	thread := frame.Thread()
	// 	frame.SetNextPC(thread.PC()) // undo anewarray
	// 	thread.InitClass(componentClass)
	// 	return
	// }

	stack := frame.OperandStack()
	count := stack.PopInt() // 第二个操作数，数组长度
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
