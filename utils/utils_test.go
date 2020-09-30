package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestItertoolsProduct(t *testing.T) {
	first := []string{"a"}
	second := []string{"d", "e"}
	result := ItertoolsProduct(first, second)
	assert.ElementsMatch(t, *result, [][2]string{[2]string{"a", "d"}, [2]string{"a", "e"}} )

}

// 写端
func write(ch chan int) {
	ch <- 100
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch <- 200
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch <- 300                      // 写入第三个，造成阻塞
	fmt.Printf("ch addr：%v\n", ch) // 没有输出
}

func TestThread(t *testing.T){

	var ch chan int        // 声明一个有缓冲的channel
	ch = make(chan int, 2) // 可以写入2个数据

	// 向协程中写入数据
	go write(ch)

	// 防止主go程提前退出，导致其他协程未完成任务
	time.Sleep(time.Second * 3)
}
