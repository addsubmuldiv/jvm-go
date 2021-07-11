package heap

import "ch07/classfile"

// MemberRef 成员变量符号引用，用于给字段、方法的符号引用嵌套的类型
type MemberRef struct {
	SymRef     // 符号引用
	name       string
	descriptor string // 这里存放了字段、方法的类型描述符，避免重名的出现，尽管java是但凡重名就不行，但是jvm是可以的，所以需要描述符进行区分
}

// 读取这个 字段、方法符号索引的所属类，字段/方法名，以及对应的描述符
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

// Name 返回字段/方法名
func (self *MemberRef) Name() string {
	return self.name
}

// Descriptor 返回字段/方法的描述符
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
