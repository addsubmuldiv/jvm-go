package rtda

import (
	"ch07/rtda/heap"
)

// Slot 局部变量表的构成元素
type Slot struct {
	num int32        // num字段存储数值类型，对于浮点数，直接把它按二进制表达的整数存
	ref *heap.Object // ref就是存储各种对象的引用
}
