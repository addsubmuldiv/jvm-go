package heap

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

// ResolvedClass 如果类D通过符号引用N引用类C的话，要解析N，先用D的类加载器加载C，
//然后检查D是否有权限访问C，如果没有，则抛出IllegalAccessError异常
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

// jvms8 5.4.3.1
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
