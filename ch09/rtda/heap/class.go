package heap

import (
	"ch09/classfile"
	"strings"
)

// Class name, superClassName and interfaceNames are all binary names(jvms8-4.2.1)
// todo 这个就是反射里面用的那个 .class ？
type Class struct {
	accessFlags       uint16
	name              string // thisClassName
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool // todo 每个类都有个常量池？	—— 是的,或者应该说每个class文件
	fields            []*Field      // 所有字段，包括 实例、静态
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint // todo 这里是说，类知道自己有几个对象的意思？	—— 这里是指 实例变量 个数
	staticSlotCount   uint
	staticVars        Slots // todo 静态变量也用 Slot ？？	—— 是的
	initStarted       bool
	jClass            *Object // java.lang.Class实例，这个和 object 对象里面的 extra 包含的指针是 反向的关系
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *Class) SuperClassName() string {
	return self.superClassName
}

func (self *Class) InterfaceNames() []string {
	return self.interfaceNames
}

func (self *Class) Fields() []*Field {
	return self.fields
}

func (self *Class) Methods() []*Method {
	return self.methods
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) Interfaces() []*Class {
	return self.interfaces
}

func (self *Class) InstanceSlotCount() uint {
	return self.instanceSlotCount
}

func (self *Class) StaticSlotCount() uint {
	return self.staticSlotCount
}

func (self *Class) Name() string {
	return self.name
}

// classfile 读取文件，然后从 classfile 里面把类相关的信息都取出来，放在这里
// 有点 java web 分层的味道
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

// jvms 5.4.4  如果类D想访问类C，需要满足两个条件之一：
// C是public，或者C和D在同一个包内。
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

// 获取一个静态方法，是通过 方法名和描述符 来操作的
func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

// 根据普通类，获取对应的数组类
func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}
func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}
func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotId, ref)
}
