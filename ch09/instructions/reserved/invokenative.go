package reserved

import (
	"ch09/instructions/base"
	"ch09/native"
	_ "ch09/native/java/lang"
	_ "ch09/native/sun/misc"
	"ch09/rtda"
)

// 0xFE 即为调用本地方法
type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor) // 查找本地方法
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
