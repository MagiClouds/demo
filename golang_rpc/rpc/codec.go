package rpc

import (
    "bytes"
    "encoding/gob"
)

//定义数据格式和编解码

//定义RPC交互的数据类型
type RPCData struct {
    Name string        `desc:"访问函数"`
    Args []interface{} `desc:"访问参数"`
}

//编码
func encode(data RPCData) ([]byte, error)  {
    var buf bytes.Buffer
    bufEnc := gob.NewEncoder(&buf)

    if err := bufEnc.Encode(data); err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}

//解码
func decode(b []byte) (RPCData, error) {
    buf := bytes.NewBuffer(b)

    bufDec := gob.NewDecoder(buf)
    var data RPCData

    if err := bufDec.Decode(&data); err != nil {
        return data, err
    }

    return data, nil
}