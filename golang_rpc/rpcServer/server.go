package main

import (
    "log"
    "net/http"
    "net/rpc"
)

type Rect struct {}

type Params struct {
    Width, Height int
}

func (r *Rect) Area(p Params, ret *int) error  {
    *ret = p.Width * p.Height
    return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error  {
    *ret = (p.Width+p.Height)*2
    return nil
}

func main()  {
    //注册
    rect := new(Rect)
    rpc.Register(rect)

    //把服务绑定到http协议上
    rpc.HandleHTTP()

    //监听服务，等待客户端调用
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
