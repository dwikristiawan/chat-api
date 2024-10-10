package ws_resource

import "github.com/gorilla/websocket"

type WsConn struct {
	Connection map[string]*websocket.Conn
}

func NewWsConn() WsConn {
	return WsConn{make(map[string]*websocket.Conn)}
}
