// 这里是俩标记属性
package classfile

// 这个表示某个方法已经不再推荐使用，未来可能移除
type DeprecatedAttribute struct {
	MarkerAttribute
}

// Synthetic属性用来标记源文件中不存在、由编译器生成的类成员，引入Synthetic属性主要是为了支持内部类和内部接口。
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
