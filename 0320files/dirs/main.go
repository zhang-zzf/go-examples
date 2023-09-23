package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var p = fmt.Println

func main() {
	createFile()
	// 在当前 wd (workingDir) 目录下创建文件夹
	err := os.Mkdir("subdir", 0755)
	if err != nil {
		p(err)
	}
	// subdir 可以是一个文件，也可以是一个目录（会删除目录下的所有文件）
	os.RemoveAll("subdir")
	// change workding dir
	err = os.Chdir("/tmp")
	if err != nil {
		p("Chdir failed ->", err)
	}
	err = os.MkdirAll("a/b/c", 0755)
	if err != nil {
		p("MkdirAll failed ->", err)
	}
	os.RemoveAll("a/b/c")
	p("TempDir =>", os.TempDir())
	p(os.UserHomeDir())
	p(os.UserCacheDir())
	p(os.UserConfigDir())
	p("----------")
	p(execFile())
	p(absPath())
	p(workingDir())
	p("----------")
	readDir("/tmp")
	readDir("/tmp/")
}
func readDir(dir string) {
	open, _ := os.Open(dir)
	entries, _ := open.ReadDir(0)
	for _, entry := range entries {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
		info, _ := entry.Info()
		fmt.Println(info.Name(), info.IsDir(), info.Mode(), info.ModTime(), info.Size())
	}

}

func execFile() string {
	executable, _ := os.Executable()
	return executable
}

func absPath() string {
	args0 := os.Args[0]
	p("args0 ->", args0)
	dir := filepath.Dir(args0)
	p("dir of args0 -> ", dir)
	absDir, _ := filepath.Abs(dir)
	return absDir
}

func workingDir() string {
	wd, _ := os.Getwd()
	return wd
}

func createFile() {
	someFile := "/tmp/a.txt"
	file, err := os.Create(someFile)
	if err != nil {
		p(err)
	}
	defer file.Close()
	file.WriteString("Hello,World.\n")
	file.Sync()
	writer := bufio.NewWriter(file)
	writer.WriteString("你好，世界。\n")
	writer.Flush()
	fileContent, err := os.ReadFile(someFile)
	if err != nil {
		p(err)
	} else {
		p(string(fileContent))
	}
	// subdir 可以是一个文件，也可以是一个空目录
	if err := os.Remove(someFile); err != nil {
		p(err)
	}
}
