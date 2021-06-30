// 因为和各种字段名、方法名本质上是一样的，都是字符串
// 所以字符串常量的就不存真正的数据，直接用一个指向UTF8常量的索引了
package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16 // 用来索引到UTF8字符串
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

// 这里能找到常量池那个对应的UTF8字符串，返回的是字符串
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)

}
