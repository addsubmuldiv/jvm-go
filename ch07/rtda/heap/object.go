// Package rtda 临时用的对象类
package heap

type Object struct {
	class  *Class
	fields Slots
}

// 就是让对象里面存个指向class的指针，然后初始化一下字段的slot
func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.fields
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}
