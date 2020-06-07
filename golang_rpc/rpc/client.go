package rpc

import (
    "net"
    "reflect"
)

type Client struct {
    conn net.Conn
}

//构造
func NewClient(conn net.Conn) *Client {
    return &Client{conn: conn}
}

func (c *Client) callRPC(rpcName string, fPtr interface{}) {
    //通过反射，获取fPtr未初始化的函数原型
    fn := reflect.ValueOf(fPtr).Elem()

    f := func(args []reflect.Value) []reflect.Value {
        inArgs := make([]interface{}, 0, len(args))
        for _, arg := range args {
            inArgs = append(inArgs, arg.Interface())
        }

        cliSession := NewSession(c.conn)

        reqRPC := RPCData{Name: rpcName, Args: inArgs}
        b, err := encode(reqRPC)
        if err != nil {
            panic(err)
        }

        //写数据
        if err := cliSession.Write(b); err != nil {
            panic(err)
        }

        //读数据
        respBytes, err := cliSession.Read()
        if err != nil {
            panic(err)
        }

        data, err := decode(respBytes)
        if err != nil {
            panic(err)
        }

        outArgs := make([]reflect.Value, 0, len(data.Args))

        for i, arg := range data.Args {
            if arg == nil {
                //必须填充一个真正的类型，不能是nil
                outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
                continue
            }
            outArgs = append(outArgs, reflect.ValueOf(arg))
        }
        return outArgs
    }

    v := reflect.MakeFunc(fn.Type(), f)
    fn.Set(v)
}
