package demo

import (
    "fmt"
    "time"
)

func ChannelDemo() {
    ch := make(chan bool)
    go func() {
        for i := 0;i < 10; i++ {
            fmt.Println(i)
            time.Sleep(time.Second)
        }
        close(ch)
    }()

    fmt.Println("Wait for sleep")
    select {
    case  <-ch:
        fmt.Println("Demo is done")
    case <-time.After(5 * time.Second):
        fmt.Println("Wait timeout")
    }
    fmt.Println("Wait done")
}
