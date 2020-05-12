

function g1Join() {
    const g1io = io('/g1');
    g1io.emit('cJoin', "room1");
}
g1Join();


// $('form').submit(
// function () {
//     chatio.emit('cbChat', $('#m').val());
// }
// );

// server.OnEvent("/chat",
// "cbChat", func(s socketio.Conn, msg string) {
//     server.BroadcastToRoom("/chat", roomName,
//         "sChat", s.ID()+" => "+msg)
// })