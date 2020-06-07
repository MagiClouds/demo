package rpc

import (
    "encoding/binary"
    "io"
    "net"
)

//回话连接的结构体
type Session struct {
    conn net.Conn
}

//构造方法
func NewSession(conn net.Conn) *Session {
    return &Session{conn: conn}
}

//向回话中写数据
func (s *Session) Write(data []byte) error {
    //4字节头记录数据长度 + 真实数据
    buf := make([]byte, 4+len(data))

    binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))

    copy(buf[4:], data)
    _, err := s.conn.Write(buf)
    if err != nil {
        return err
    }
    return nil
}

//从连接中读数据
func (s *Session) Read() ([]byte, error) {
    //读头部长度
    header := make([]byte, 4)

    if _, err := io.ReadFull(s.conn, header); err != nil {
        return nil, err
    }

    dataLen := binary.BigEndian.Uint32(header)

    data := make([]byte, dataLen)
    _, err := io.ReadFull(s.conn, data)
    if err != nil {
        return nil, err
    }

    return data, nil
}
