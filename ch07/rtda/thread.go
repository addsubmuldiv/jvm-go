package rtda

import "ch07/rtda/heap"

// Thread 运行时数据区的半壁江山，线程
// 每个线程有对应的程序计数器，以及一个jvm虚拟机栈
// 栈里面存储的是帧Frame
// 所以说，一个线程，除了一个程序计数器PC，就是一堆以栈的形式存起来的帧
// 而帧又是  局部变量表 + 操作数栈，以及一个指向下一条指令的PC
type Thread struct {
	pc    int    // 程序计数器
	stack *Stack // jvm虚拟机栈
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// CurrentFrame 获取当前栈顶的帧，即当前帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

// NewFrame 新建一个栈帧，俩参数是最大局部变量表大小、最大操作数栈大小
func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
