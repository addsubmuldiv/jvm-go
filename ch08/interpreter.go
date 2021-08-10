package main

import (
	"ch08/instructions"
	"ch08/instructions/base"
	"ch08/rtda"
	"ch08/rtda/heap"
	"fmt"
)

// 虚拟机执行代码的起始入口，从 class 文件里面的 methodInfo 入手
func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread() // 新建一个线程
	frame := thread.NewFrame(method)
	thread.PushFrame(frame) // 线程本质就是一堆以栈形式存起来的栈帧 + 一个PC，这里就是新建一个帧然后放进虚拟机栈
	defer catchErr(frame)
	loop(thread, logInst)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("")
	}
}

// 这个就是一个最主要的函数，无限循环地
// 在一个方法里面读取指令，执行，但是貌似没有看到如何执行函数？
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
