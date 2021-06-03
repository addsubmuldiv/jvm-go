package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16 // 指向类名的UTF8字符串
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
