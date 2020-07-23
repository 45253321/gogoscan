package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrFileNotFound = errors.New("The file not found")

func ReadFile(path string) (*[]string, error){

	content := make([]string, 10, 10)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("read file err:", err)
		return nil, nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, _, err := reader.ReadLine()
		if err != nil && err != io.EOF{
			fmt.Println("read err:", err)
			break
		}
		content = append(content, string(str))
		if err == io.EOF {
			break
		}
	}
	return &content, nil
}

func main()  {
	contents, err := ReadFile("./password_ssh.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	for _, line := range *contents {
		fmt.Printf(line)
	}

}