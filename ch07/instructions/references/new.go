package references

import (
	"ch07/instructions/base"
	"ch07/rtda"
	"ch07/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

// 先解析出class，然后调用NewObject
func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	// todo: init class

	if class.IsInterface() || class.IsAbstract() { // 抽象类和接口都不能实例化
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
