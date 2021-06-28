package rtda

import "ch07/rtda/heap"

// Frame 栈帧，由于虚拟机栈采用的链栈实现
// 因此有个lower指向栈中的下一个元素
// 每个栈帧都有个局部变量表和操作数栈
// 每个栈帧也都知道自己是哪个线程的
// todo 以及一个指向下一条指令的PC —— 这个是由 reader 里面控制的，每从字节码读取一点，就增加
type Frame struct {
	lower        *Frame // stack is implemented as linked list
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int // the next instruction after the call
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// getters & setters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) Method() *heap.Method {
	return self.method
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
