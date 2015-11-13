// Package main provides io test
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileInfo interface {
	Name() string
	Size() int64
	IsDir() bool
	Sys() interface{}
}

func ReadDirTest() {
	// 读取目录 ioutil.ReadDir
	dir_list, err := ioutil.ReadDir("./log")
	if err != nil {
		fmt.Println("read dir error")
		return
	}
	for i, v := range dir_list {
		fmt.Println(i, "=", v.Name())
		fmt.Println(v.Name(), "的文件大小：", v.Size())
		fmt.Println(v.Name(), "的创建时间：", v.ModTime())
		fmt.Println(v.Name(), "的系统信息：", v.Sys())
		if v.IsDir() == true {
			fmt.Println(v.Name(), "是目录")
		}
	}
}
func ReadFileTest() {
	data, err := ioutil.ReadFile("./log/1.log")
	if err != nil {
		fmt.Println("read error")
		os.Exit(1)
	}
	fmt.Println(string(data))
}
func WriteFile(logdata string) {
	file, err := ioutil.TempFile("./log", "reallog.md")
	defer file.Close()
	if err != nil {
		fmt.Println("Failed to create file")
		return
	}
	file.WriteString("Log:" + logdata)
	filedata, _ := ioutil.ReadFile(file.Name())
	fmt.Println(string(filedata))
}
func main() {
	ReadDirTest()
	ReadFileTest()
	WriteFile("gogolog")

}
