// 方法或者字段的描述符
package classfile

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16	// 方法或者字段名的名字
	descriptorIndex uint16	// 字段描述符/方法描述符，就是表示啥类型、返回啥类型、参数啥类型，用于支持重载
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
