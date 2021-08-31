package references

import (
	"ch11/instructions/base"
	"ch11/rtda"
	"ch11/rtda/heap"
)

// GET_FIELD 根据自己的索引从常量池获取对应的符号引用
type GET_FIELD struct {
	base.Index16Instruction
}

// Execute 通过字段符号引用获取字段的实际对象，能调用到这里的一定是对象而不是普通类型了
func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	stack := frame.OperandStack()
	ref := stack.PopRef() // 这个实现有一个假设，就是你在获取这个字段的时候，这个字段所属的那个类的引用一定在操作数栈栈顶
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}
}
