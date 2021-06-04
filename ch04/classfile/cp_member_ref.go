// 类里面的字段、方法对应的类名和名字、描述符
// class_index就是这个字段、方法所属的类名，name_and_type就是字段名、描述符(类型、返回值之类)
// 注意！！ 这里的都是符号引用，是编译原理的东西，主要是为了支持动态链接操作
// 和java类里面定义的字段、方法不一样
package classfile

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16 // 常量池索引，指向类型
	nameAndTypeIndex uint16 // 常量池索引，指向这个字段或者方法的描述符
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// 嵌套继承
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }
