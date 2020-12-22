package main

import "fmt"
import "time"

func main(){
    // 定时器表示在未来某一时刻的独立事件。你告诉
    //定时器需要等待的时间，然后他将提供一个用于
    // 通知的通道，这里的定时器将等待两秒
    timer1 := time.NewTimer(time.Second * 2)

    // <- timer1,C 知道这个定时器的通道C明确发送了
    // 定时器失效的值之前，将一直阻塞
    <-timer1.C
    fmt.Println("Timer 1 expired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()

    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}

