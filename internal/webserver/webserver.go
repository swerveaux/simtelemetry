package webserver

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
)

type SocketCollection struct {
	Sockets    map[string]chan []byte
	IncomingCh chan []byte
	mu         sync.Mutex
}

func (s *SocketCollection) AddSocket(id string, ch chan []byte) {
	s.mu.Lock()
	s.Sockets[id] = ch
	s.mu.Unlock()
}

func (s *SocketCollection) RemoveSocket(id string) {
	s.mu.Lock()
	socket := s.Sockets[id]
	delete(s.Sockets, id)
	close(socket)
	s.mu.Unlock()
}

func NewSocketCollection(ch chan []byte) *SocketCollection {
	s := SocketCollection{IncomingCh: ch}
	s.Sockets = make(map[string]chan []byte)
	var rpm Telemetry
	go func() {
		for msg := range ch {
			r := bytes.NewReader(msg)
			err := binary.Read(r, binary.LittleEndian, &rpm)
			if err != nil {
				fmt.Println(err)
				return
			}
			j, err := json.Marshal(rpm)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, socket := range s.Sockets {
				socket <- j
			}
		}
	}()
	return &s
}

func Run(port int, ch chan []byte) error {
	socketCollection := NewSocketCollection(ch)
	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/websocket", websocketHandler(socketCollection))

	http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
	return nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketHandler(s *SocketCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan []byte)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		u, err := uuid.NewV4()
		if err != nil {
			log.Println(err)
			return
		}
		s.AddSocket(u.String(), ch)
		for msg := range ch {

			err = conn.WriteMessage(websocket.TextMessage, msg)

			if err != nil {
				s.RemoveSocket(u.String())
				return
			}
		}
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	page := `<html>
<head>
<title>This is a test</title>
<script type="text/javascript">
window.addEventListener("DOMContentLoaded", function() {
var exampleSocket = new WebSocket("ws://192.168.86.35:10001/websocket");
var gearfield = document.getElementById("gear");
var rpmfield = document.getElementById("rpm");
var accelfield = document.getElementById("accel");
exampleSocket.onmessage = function(event) {
  var d = JSON.parse(event.data);
  gearfield.innerText = d["gear"]
  rpmfield.innerText = d["current_engine_rpm"];
  accelfield.innerText = d["accel"];
}
}, false);
</script>
</head>
<body>
<h1>Whassup?</h1><br />
<table border="0">
  <tbody>
    <tr>
      <td>Gear</td><td id="gear">
    </tr>
    <tr>
      <td>RPM</td><td id="rpm">
    </tr>
    <tr>
      <td>Accel</td><td id="accel">
    </tr>
  </tbody>
</table>
</body>
</html>`
	fmt.Fprintf(w, page)
}
