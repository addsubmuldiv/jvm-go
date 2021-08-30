package rtda

// Stack 虚拟机栈使用链表实现
// 链栈采用头插法实现
type Stack struct {
	maxSize uint   // 虚拟机栈最大容量
	size    uint   // 当前栈中有几个Frame
	_top    *Frame // 栈顶元素
}

// 栈的初始化，参数指定虚拟机栈最大容量
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// 压入栈帧
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if self._top != nil {
		frame.lower = self._top //插入头部
	}

	self._top = frame // 更新栈顶
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--

	return top
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	return self._top
}

func (self *Stack) isEmpty() bool {
	return self._top == nil
}

func (self *Stack) clear() {
	for !self.isEmpty() {
		self.pop()
	}
}

func (self *Stack) getFrames() []*Frame {
	frames := make([]*Frame, 0, self.size)
	for frame := self._top; frame != nil; frame = frame.lower {
		frames = append(frames, frame)
	}
	return frames
}
