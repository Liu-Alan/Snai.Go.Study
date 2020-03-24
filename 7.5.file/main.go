package main

import (
	"fmt"
	"os"
)

//目录操作
/*
func main() {
	os.Mkdir("alan", 0777)
	os.MkdirAll("alan/test1/test2", 0777)

	err := os.Remove("alan")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("alan")
}
*/
//写文件
/*
func main() {
	fileName := "alan.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(fileName, err)
		return
	}
	defer file.Close()
	for i := 0; i < 10; i++ {
		file.WriteString("Just a test(string)!\r\n")
		file.Write([]byte("Just a test(byte)!\r\n"))
	}
}
*/
//读文件
/*
func main() {
	fileName := "alan.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fileName, err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}

		os.Stdout.Write(buf[:n])
	}
}
*/
//删除文件
func main() {
	fileName := "alan.txt"
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(fileName, err)
		return
	}
}
