package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

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

type userinfo struct {
	minute string
	second string
}

func (u *userinfo) GetStrTime() string {
	return u.minute + ":" + u.second
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	roomName := "Room1"
	server.OnEvent("/chat",
		"cJoin", func(s socketio.Conn) {
			fmt.Println("/chat :", s.ID()) //
			user := &userinfo{minute: strconv.Itoa(time.Now().Minute()),
				second: strconv.Itoa(time.Now().Second())}
			s.SetContext(user)
			server.JoinRoom("/chat", roomName, s)
			server.BroadcastToRoom("/chat", roomName,
				"sChat", "["+user.GetStrTime()+"]"+"Join id: "+s.ID()+", room: "+roomName)
		})

	server.OnEvent("/chat",
		"cbChat", func(s socketio.Conn, msg string) {
			server.BroadcastToRoom("/chat", roomName,
				"sChat", s.ID()+" => "+msg)
		})

	server.OnEvent("/chat",
		"cExit", func(s socketio.Conn) {
			fmt.Println("[ EXIT ID ]: ", s.ID())
			s.Close()
			server.BroadcastToRoom("/chat", roomName,
				"sExit", "[ Exit ID ]: "+s.ID())
		})

	server.OnError("/chat", func(s socketio.Conn, e error) {
		fmt.Println("[ ERROR ]: ", e, " [ ID ]: ", s.ID())
	})
	server.OnDisconnect("/chat", func(s socketio.Conn, reason string) {
		fmt.Println("[ Disconnect ]: ", reason, " [ ID ]: ", s.ID())
	})

	go server.Serve()
	defer server.Close()

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("/ :", s.ID()) //
		return nil
	})
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/socket.io/", server)

	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
