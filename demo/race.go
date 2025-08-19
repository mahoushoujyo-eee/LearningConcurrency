package demo

import "fmt"


func Race() {
    var data int
    go func() {
        data++
    }()

    if data == 0 {
        fmt.Printf("the value is %d\n", data)
    }else{
		fmt.Printf("the value is %d\n", data)
	}
}