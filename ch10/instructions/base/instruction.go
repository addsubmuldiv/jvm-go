package base

// 这个文件里面主要是对指令接口的定义，以及定义了一些可以用于嵌套复用的指令
import "ch10/rtda"

type Instruction interface {
	// FetchOperands 获取操作数的这个方法，使用 BytecodeReader 作为参数，调用的这里面的读取不同字节数的方法来读取操作数，也就是说这个 reader 的 PC 会自动更新的
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// BranchInstruction 用于给分支指令嵌套用
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// Index8Instruction 用于给读取一个字节操作数的那些指令嵌套用
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// Index16Instruction 用于给读取两个字节操作数的那些指令嵌套用
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
