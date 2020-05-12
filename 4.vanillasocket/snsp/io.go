package snsp

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/stacew/gowgtest/4.vanillasocket/snsp/g/g1"
	"github.com/stacew/gowgtest/4.vanillasocket/snsp/g/g2"
)

const (
	nsp = "/"
)

func onConnect(c socketio.Conn) error {
	log.Println(nsp, "[ Con ]ID:"+c.ID(), "R_Addr:"+c.RemoteAddr().String())

	c.Emit("sid", c.ID())
	return nil
}
func onError(c socketio.Conn, e error) {
	log.Println(nsp, "[ ERR ]ID:"+c.ID(), e)
}
func onDisconnect(c socketio.Conn, reason string) {
	log.Println(nsp, "[ Dis ]ID:"+c.ID(), reason)
}

//NewServerAndBaseHandler : New Socket Server && Base Handler
func NewServerAndBaseHandler() (*socketio.Server, error) {
	s, err := socketio.NewServer(nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	s.OnConnect(nsp, onConnect)
	s.OnError(nsp, onError)
	s.OnDisconnect(nsp, onDisconnect)

	g1.MakeHandler(s)
	g2.MakeHandler(s)

	return s, nil
}

//-ser //////////////////////////////
//NewServer(c *engineio.Options) (*Server, error)
//Close() error
//OnConnect(nsp string, f func(Conn) error)
//OnDisconnect(nsp string, f func(Conn, string))
//OnError(nsp string, f func(Conn, error))
//OnEvent(nsp, event string, f interface{})
//nsp에 사용자 이벤트 함수 등록
//ServeHTTP(w http.ResponseWriter, r *http.Request)
//???

//-ser.broad //////////////////////////////
//JoinRoom(namespace, room string, connection Conn) bool
//nsp에 방과 연결 추가
//LeaveRoom(namespace, room string, connection Conn) bool
//nsp에 방과 연결 제거
//LeaveAllRooms(namespace string, connection Conn) bool
//모든 방 및 연결 제거
//ClearRoom(namespace, room string) bool
//해당 nsp에 방 정리
//BroadcastToRoom(namespace, room, event string, args ...interface{}) bool
//해당 nsp에 방의 모두에게 이벤트와 인자를
//RoomLen(namespace, room string) int
//해당 nsp 방의 con 수
//Rooms(namespace string) []string
//해당 nsp에 모든 방 목록
//ForEach(namespace, room string, f EachFunc) bool
//방의 모두에게 특정 함수를 실행

//--con //////////////////////////////
//Close() error
//특정 con 연결 해제
//Emit(msg string, v ...interface{})
//특정 con 메시지

//Context() interface{}
//SetContext(v interface{})
//con 확장

//con 강제퇴장..?..?..
//Join(room string)
//Leave(room string)
//LeaveAll()
//Rooms() []string
