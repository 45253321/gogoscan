package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"errors"
)

var ErrFileNotFound = errors.New("The file not found")

// ---------------------------------------------------------
// 读取文件
// ---------------------------------------------------------
func readFile(path string) (*[]string, error){

	content := make([]string, 0, 0)

	file, err := os.Open(path)
	if err != nil {
		//fmt.Println("read file err:", err)
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, _, err := reader.ReadLine()
		if err != nil && err != io.EOF{
			fmt.Println("read err:", err)
			break
		}
		if err == io.EOF {
			break
		}
		content = append(content, string(str))
	}
	return &content, nil
}

// ---------------------------------------------------------
// 读取行
// ---------------------------------------------------------
func ReadLines(path string) *[]string {
	contents, err := readFile(path)
	if err != nil{
		panic(ErrFileNotFound)
	}

	return contents
}