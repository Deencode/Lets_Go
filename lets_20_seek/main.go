package main

import (
	"fmt"
	"io"
	"os"
)

const (
	//123567890
	ROOT     = "D:/Go_workspace/src/Lets_Go/lets_20_seek/"
	SrcPath  = ROOT + "seek.txt"
	tempFile = ROOT + "t.txt"
)

func main() {
	file, _ := os.Open(SrcPath)
	defer file.Close()
	con := make([]byte, 3) //每次只读3给字节
	file.Read(con)
	fmt.Println(string(con))
	//tf := getTempFile()
	con = append(con, []byte("这是插入的内容")...) //把之前源文件读取的内加上需要插入的内容
	//tf.WriteString(string(con))
	//defer tf.Close()
	file.Seek(3, 0)
	v := [1]byte{} //每次读入的内容存在这里
	for {
		_, err := file.Read(v[:])
		con = append(con, v[:]...) //将剩下读取写到已经拼接好的切片中
		if err == io.EOF {
			break
		}
	}
	//然后把拼接完的切片在存储的源文件中
	srcFile := getSrcFile()
	defer srcFile.Close()
	srcFile.WriteString(string(con))
	fmt.Println("edit src file succeed.")
}
func test1() {
	file, _ := os.Open(SrcPath) //源文件
	defer file.Close()
	file.Seek(0, 0)  //移动光标
	con := [1]byte{} //一次读取多少给字节
	file.Read(con[:])
	fmt.Println(string(con[:]))
}

func getSrcFile() os.File {
	file, _ := os.OpenFile(SrcPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	return *file
}
