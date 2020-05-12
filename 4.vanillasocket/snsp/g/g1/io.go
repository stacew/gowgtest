package g1

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

const (
	nsp = "/g1"
)

type g1User struct {
	room string
	cID  string
	x    int
	y    int
}

func onConnect(c socketio.Conn) error {
	log.Println(nsp, "[ Con ]ID:"+c.ID(), "R_Addr:"+c.RemoteAddr().String())
	return nil
}
func onError(c socketio.Conn, e error) {
	log.Println(nsp, "[ ERR ]ID:"+c.ID(), e)
}
func onDisconnect(c socketio.Conn, reason string) {
	log.Println(nsp, "[ Dis ]ID:"+c.ID(), reason)
}

func cJoin(c socketio.Conn, joinRoom string) {
	log.Println(nsp, "[ Test cJoin ]ID:"+c.ID())
	c.SetContext(&g1User{room: joinRoom, cID: c.ID(), x: 0, y: 0})
	c.Join(joinRoom)
	log.Println(c.Rooms())
}
func cLeave(c socketio.Conn) {
	log.Println(nsp, "cLeave:"+c.ID())
	user := c.Context().(*g1User)
	c.Leave(user.room)
}

func cInputSpace(c socketio.Conn) {
	log.Println(nsp, "[TEST]", "cInputSpace:"+c.ID())

}

func MakeHandler(s *socketio.Server) {

	s.OnConnect(nsp, onConnect)
	s.OnError(nsp, onError)
	s.OnDisconnect(nsp, onDisconnect)

	s.OnEvent(nsp, "cJoin", cJoin)
	s.OnEvent(nsp, "cLeave", cLeave)
	s.OnEvent(nsp, "cInputSpace", cInputSpace)
}
