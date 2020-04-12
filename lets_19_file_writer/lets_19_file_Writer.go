package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	filePath = "D:/Go_workspace/src/Lets_Go/lets_19_file_writer/test.txt"
)

//go语言中文件写入操作
func main() {
	wr1()
	wr2()
	wr3()
}
func wr1() {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("文件打开失败~~")
	}
	defer file.Close()
	file.Write([]byte("岳宁宁在Google云计算部门,她负责的是Google Cloud CDN产品!"))
	file.WriteString("\n而我却是给菜鸟~~sorry!")
}
func wr2() {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("文件打开失败~~")
	}
	defer file.Close()
	wr := bufio.NewWriter(file)
	wr.WriteString("\nYNN哈哈哈哈")
	wr.Flush() //bufio在缓存里面 必须调用这个
}
func wr3() {
	content := "不知道会不会有人看到这个内容哈哈哈哈哈~^_^"
	ioutil.WriteFile(filePath, []byte(content), 0777)
}
