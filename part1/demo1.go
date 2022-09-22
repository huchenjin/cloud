package part1

import (
	"fmt"
	"time"
)

/**
给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
*/

var strArr = [...]string{"I", "am", "stupid", "and", "weak"}

var strArr2 = [...]string{"I", "am", "smart", "and", "strong"}

func AlertArray() {
	for i, _ := range strArr {
		if i == 2 || i == 4 {
			strArr[i] = strArr2[i]
		}
	}

	fmt.Println(strArr)
}

/**
队列：
队列长度 10，队列元素类型为 int
生产者：
每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/

func Producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
		fmt.Printf("produce data:%d\n", i)
	}
	close(ch)

}

func Consumer(ch chan int) {
	for v := range ch {
		fmt.Printf("the data is %d", v)
	}
}
