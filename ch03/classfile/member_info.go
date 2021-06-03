package classfile

// 类自己定义的成员变量和方法都用这个结构体描述
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16          // 访问标记
	nameIndex       uint16          // 名字
	descriptorIndex uint16          // 描述符
	attributes      []AttributeInfo // 如果是成员变量，那这里会有个ConstantValue属性，指向这个字段的值在常量池的索引
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

// 从常量池查找字段名/方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// TODO 这是个啥
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
