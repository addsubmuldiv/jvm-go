package main

import (
	"ch05/classfile"
	"ch05/instructions"
	"ch05/instructions/base"
	"ch05/rtda"
	"fmt"
)

// 虚拟机执行代码的起始入口，从 class 文件里面的 methodInfo 入手
func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute() // 获取这个方法的代码属性
	maxLocals := codeAttr.MaxLocals()      // 获取已经计算好的局部变量表
	maxStack := codeAttr.MaxStack()        // 最大操作数栈
	bytecode := codeAttr.Code()            // 获取字节码

	thread := rtda.NewThread() // 新建一个线程
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame) // 线程本质就是一堆以栈形式存起来的栈帧 + 一个PC，这里就是新建一个帧然后放进虚拟机栈

	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("")
	}
}

// 这个就是一个最主要的函数，无限循环地
// 在一个方法里面读取指令，执行，但是貌似没有看到如何执行函数？
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC() // 这里获取了，然后下面又更新了？TODO 什么意思
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()                // java 用一个字节来编码指令， 所以读取一个字节
		inst := instructions.NewInstruction(opcode) // 生成指令对象
		inst.FetchOperands(reader)                  // 根据指令的不同实现，获取各自操作数
		frame.SetNextPC(reader.PC())                // 把当前栈帧的 PC 更新

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
