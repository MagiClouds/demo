package main

import (
    "fmt"
    "log"
    "net/rpc"
)

type Params struct {
    Width, Height int
}

func main() {
    //连接
    client, err := rpc.DialHTTP("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }

    //调用服务端方法
    var response int
    if err := client.Call("Rect.Area", Params{Width: 20, Height: 30}, &response); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Area is ", response)

    if err := client.Call("Rect.Perimeter", Params{10, 20}, &response); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Perimeter is ", response)
}
