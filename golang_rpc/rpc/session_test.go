package rpc

import (
    "fmt"
    "log"
    "net"
    "sync"
    "testing"
)

func TestSession(t *testing.T) {
    addr := "localhost:8000"
    data := "hello world"

    wg := sync.WaitGroup{}

    wg.Add(2)

    go func() {
        defer wg.Done()

        lis, err := net.Listen("tcp", addr)
        if err != nil {
            log.Fatal(err)
        }

        conn, err := lis.Accept() // 监听
        if err != nil {
            log.Fatal(err)
        }

        s := NewSession(conn)

        if err := s.Write([]byte(data)); err != nil {
            log.Fatal(err)
        }
    }()

    go func() {

        defer wg.Done()

        conn, err := net.Dial("tcp", addr)
        if err != nil {
            log.Fatal(err)
        }

        s := NewSession(conn)

        b, err := s.Read()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(string(b))
    }()
    wg.Wait()
}
