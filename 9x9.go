/* 九九口诀表

1. 加法口诀表
2. 减法口诀表
*/

package main

import (
    "fmt"
)

// 加法口诀表
func add9x9() {
    fmt.Println("加法口诀表")
    for i := 1; i <= 9; i++ {
        for j := 1; j <= 9; j++ {
            if j < i {
                fmt.Print("       ")
                continue
            }
            if i + j < 10 {
                fmt.Printf("%d+%d=%d  ", i, j, i+j)
            } else {
                fmt.Printf("%d+%d=%d ", i, j, i+j)
            }
        }
        fmt.Println()
    }
}

//减法口决表
func sub9x9() {
    fmt.Println("减法口决表")
    for i := 1; i <= 9; i++ {
        for j := 1; j <= 9; j++ {
            if j > i {
                fmt.Print("       ")
                continue
            }
            fmt.Printf("%d-%d=%d  ", i, j, i-j)
        }
        fmt.Println()
    }
}

func main() {
    add9x9()
    fmt.Println()
    sub9x9()
}