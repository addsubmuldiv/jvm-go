package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	cp ConstantPool
	stringIndex uint16	// 用来索引到UTF8字符串
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

// 这里能找到那个对应的UTF8字符串
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)

}
