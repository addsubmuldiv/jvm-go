package heap

// 符号引用公用的嵌套类
type SymRef struct {
	cp        *ConstantPool // 存放符号引用所在常量池的指针
	className string        // 符号引用的类名
	class     *Class        // 缓存被引用的那个类型的指针
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

// resolveClassRef 如果类D通过符号引用N引用类C的话，要解析N，先用D的类加载器加载C，
//然后检查D是否有权限访问C，如果没有，则抛出IllegalAccessError异常
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
