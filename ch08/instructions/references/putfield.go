package references

import (
	"ch08/instructions/base"
	"ch08/rtda"
	"ch08/rtda/heap"
)

// PUT_FIELD Set field in object
// 三个操作数，第一个是符号引用的常量池索引，第二个是值，第三个是实例对象的引用
// 值是从当前帧的操作数栈里面取
type PUT_FIELD struct{ base.Index16Instruction } // 这个是 符号引用 在常量池的索引

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef) // todo 为什么对象的字段的符号引用也在常量池？ ——常量池是和classfile绑定的貌似，所以会有这种，反正用到的各种东西就忘常量池放
	field := fieldRef.ResolvedField()                       // todo 一个java文件里面定义的多个类，常量池一样？？ —— 这个接受的都是class文件，应该是单独一个类一个class文件的，即使一开始在一个java文件里面

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()

	// 以下，根据 field 的不同类型，从操作数栈里面弹一个值出来，给它赋值
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	default:
		// todo
	}
}
