package heap

// jvm里面的描述符都是一些莫名其妙的字符串，需要解析出来
type MethodDescriptor struct {
	parameterTypes []string // 方法的参数类型
	returnType     string   // 方法返回类型
}

func (self *MethodDescriptor) addParameterType(t string) {
	pLen := len(self.parameterTypes)
	if pLen == cap(self.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, self.parameterTypes)
		self.parameterTypes = s
	}

	self.parameterTypes = append(self.parameterTypes, t)
}
