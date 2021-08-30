package heap

import "ch10/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

// 从classfile类构建class结构的时候用
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 符号引用解析
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

// jvms 5.4.3.2
// 解析字段符号引用A，先解析这个字段所属的类B的符号引用，然后从这个类B里面查找这个字段A是否存在，
// 再判断使用这个符号引用的类C是否有对A的访问权限
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.field = field
}

// 从类中查找字段
func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields { // 先自己
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range c.interfaces { //再接口
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil { // 再父类
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
