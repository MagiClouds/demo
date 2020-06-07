package rpc

import (
    "fmt"
    "net"
    "reflect"
)

//声明服务端
type Server struct {
    addr string
    //服务端维护的函数名到函数反射值的map
    funcs map[string]reflect.Value
}

//构造
func NewServer(addr string) *Server {
    return &Server{
        addr:  addr,
        funcs: make(map[string]reflect.Value),
    }
}

//服务注册
func (s *Server) Register(rpcName string, f interface{}) {
    if _, ok := s.funcs[rpcName]; ok {
        return
    }

    fVal := reflect.ValueOf(f)
    s.funcs[rpcName] = fVal
}

func (s *Server) Run() {
    lis, err := net.Listen("tcp", s.addr)
    if err != nil {
        fmt.Printf("listen %s error: %v", s.addr, err)
        return
    }

    for {
        conn, err := lis.Accept()
        if err != nil {
            fmt.Printf("accept error: %v", err)
            return
        }

        session := NewSession(conn)
        b, err := session.Read()
        if err != nil {
            fmt.Printf("read error: %v", err)
            return
        }

        rpcData, err := decode(b)
        if err != nil {
            fmt.Printf("decode error: %v", err)
            return
        }
        f, ok := s.funcs[rpcData.Name]
        if !ok {
            fmt.Printf("function %s is not existe", rpcData.Name)
            break
        }

        args := make([]reflect.Value, 0, len(rpcData.Args))

        for _, v := range rpcData.Args {
            args = append(args, reflect.ValueOf(v))
        }

        out := f.Call(args)

        outArgs := make([]interface{}, 0, len(out))
        for _, v := range out {
            outArgs = append(outArgs, v.Interface())
        }

        respRpcData := RPCData{Name: rpcData.Name, Args: outArgs}
        respBytes, err := encode(respRpcData)
        if err != nil {
            fmt.Printf("encode error: %v", err)
            return
        }

        if err := session.Write(respBytes); err != nil {
            fmt.Printf("session write data error: %s", err)
        }

    }

}
