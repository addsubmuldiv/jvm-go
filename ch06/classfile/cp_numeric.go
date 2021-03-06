// 4个数字常量及其对应的readInfo方法
package classfile

import (
	"math"
)

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes) // 读取一个uint32，转化成float32
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
