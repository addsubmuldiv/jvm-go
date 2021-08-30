package heap

import "ch10/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.3
// 这里是 d 调用 c 的方法  (非接口方法)
// 目的是根据 方法的符号引用 获取 方法本身
func (self *MethodRef) resolveMethodRef() {
	//class := self.Class()
	d := self.cp.class
	c := self.ResolvedClass() // 根据  方法的符号引用self  解析出定义这个方法的类 c
	if c.IsInterface() {      // 这里指的是， 非接口方法
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

// 现在 c 的类里面找，找不到就去 c 实现的接口里面找
func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
