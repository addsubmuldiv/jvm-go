package rtda

import (
	"ch11/rtda/heap"
	"math"
)

// LocalVars 主要用于存储方法参数（形参）和定义在方法体内的局部变量，最基本的存储单位是Slot。
// 存放编译期可知的各种基本数据类型，8种数据类型，引用类型（reference），returnAddress类型的变量。
// 32位以内的类型占用一个Slot，包括returnAddress类型，64位的类型（long和double）占用2个Slot。
// byte、short、char在存储之前转换为int，boolean在存储前转换为int，0标识false，1 标识true。
type LocalVars []Slot

// 以下就是各类数据类型的get和set方法，需要指定索引
// TODO 但是这个索引是哪里来的？	—— 看 iload.go

// 局部变量表的大小是编译器一开始算好的
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

// TODO 如何知道哪个下标对应的是哪个类型的数据？	—— 读取的时候，有些是指令隐含的，有些是指令读取了一个字节的操作数，由这个操作数指定，在 iload.go 中解答了
func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}
func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}

func (self LocalVars) GetThis() *heap.Object {
	return self.GetRef(0)
}

func (self LocalVars) SetSlot(index uint, slot Slot) {
	self[index] = slot
}
