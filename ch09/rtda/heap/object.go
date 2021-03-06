// Package rtda 临时用的对象类
package heap

type Object struct {
	class *Class
	//fields Slots
	data  interface{}
	extra interface{} // extra字段用来记录Object结构体实例的额外信息, 第9章，只用它来记录类对象对应的Class结构体指针。
}

func (self *Object) Extra() interface{} {
	return self.extra
}

func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

// 就是让对象里面存个指向class的指针，然后初始化一下字段的slot
// 仅针对普通对象
func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

// reflection
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
