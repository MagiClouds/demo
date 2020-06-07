package rpc

import (
    "encoding/gob"
    "fmt"
    "net"
    "testing"
)

type User struct {
    Name string
    Age  int
}

func queryUser(id int) (User, error)  {
    user := make(map[int]User)

    user[1] = User{"aa", 25}
    user[2] = User{"bb", 26}
    user[3] = User{"cc", 27}
    user[4] = User{"dd", 28}

    if u, ok := user[id]; ok {
        return u, nil
    }

    return User{}, fmt.Errorf("id[%d] not in user db", id)
}

func TestRPC(t *testing.T)  {
    gob.Register(User{})


    addr := "localhost:8080"

    srv := NewServer(addr)
    srv.Register("queryUser", queryUser)

    go srv.Run()

    conn, err := net.Dial("tcp", addr)
    if err != nil {
        t.Error(err)
    }

    cli := NewClient(conn)
    var query func(int)(User, error)
    cli.callRPC("queryUser", &query)
    user, err := query(1)
    if err != nil {
        t.Error(err)
    }
    fmt.Println(user.Name, user.Age)
}




