package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// 类似于 给定 a/b/c/*
// 这里直接扫一遍给定base目录，base目录里面的目录直接略过，这里只读取jar
// todo 那么zip呢？？
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}
