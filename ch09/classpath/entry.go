//@Title  entry.go
//@Description  定义类路径
//@Update  ${DATE} ${TIME}
package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator) //路径分隔符，默认为 ';'

//类路径接口
type Entry interface {
	readClass(className string) ([]byte, Entry, error) // 在对应的类路径下，根据类名把类的字节码读取出来
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {

		return newZipEntry(path)
	}

	return newDirEntry(path)
}
