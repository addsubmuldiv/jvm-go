package references

import (
	"ch10/instructions/base"
	"ch10/rtda"
	"ch10/rtda/heap"
)

const (
	AT_BOOLEAN = iota + 4
	AT_CHAR
	AT_FLOAT
	AT_DOUBLE
	AT_BYTE
	AT_SHORT
	AT_INT
	AT_LONG
)

// 俩操作数，第一个atype
// 创建基本类型数组
type NEW_ARRAY struct {
	atype uint8 // 表示要创建哪种类型的数组
}

func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUint8()
}

func (self *NEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt() // 该指令的第二个操作数，数组长度
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, self.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

// 调用类加载器，获取到特制的数组类
func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}
