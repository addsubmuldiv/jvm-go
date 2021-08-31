package base

import (
	"ch11/rtda"
	"ch11/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot() // 从 invokerFrame 的操作数栈里面取出参数
			newFrame.LocalVars().SetSlot(uint(i), slot)   // 放到 被执行方法的那个 newFrame 的局部变量表里面
		}
	}

	// hack!  暂时略过本地方法
	//if method.IsNative() {
	//	if method.Name() == "registerNatives" {
	//		thread.PopFrame()
	//	} else {
	//		panic(fmt.Sprintf("native method: %v.%v%v\n",
	//			method.Class().Name(), method.Name(), method.Descriptor()))
	//	}
	//}
}
