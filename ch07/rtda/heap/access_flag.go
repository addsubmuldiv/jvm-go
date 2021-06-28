package heap

const (
	ACC_PUBLIC       = 0x0001 // class field method		公有
	ACC_PRIVATE      = 0x0002 //       field method		私有
	ACC_PROTECTED    = 0x0004 //       field method		只允许自己和子类访问
	ACC_STATIC       = 0x0008 //       field method		静态
	ACC_FINAL        = 0x0010 // class field method		不可变
	ACC_SUPER        = 0x0020 // class					超类
	ACC_SYNCHRONIZED = 0x0020 //             method		同步
	ACC_VOLATILE     = 0x0040 //       field			线程共享变量随时写回主存
	ACC_BRIDGE       = 0x0040 //             method
	ACC_TRANSIENT    = 0x0080 //       field			序列化时进行忽略
	ACC_VARARGS      = 0x0080 //             method
	ACC_NATIVE       = 0x0100 //             method
	ACC_INTERFACE    = 0x0200 // class					接口
	ACC_ABSTRACT     = 0x0400 // class       method		抽象
	ACC_STRICT       = 0x0800 //             method
	ACC_SYNTHETIC    = 0x1000 // class field method		编译器自己添加的字段、方法、类
	ACC_ANNOTATION   = 0x2000 // class					注解
	ACC_ENUM         = 0x4000 // class field			枚举
)
